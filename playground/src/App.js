import './App.css';
import React, { useState, useEffect } from 'react';
import {
  Box,
  TextField,
  Button,
  Select,
  MenuItem,
  FormControl,
  Typography,
  InputLabel,
  OutlinedInput,
  Chip,
  Tabs,
  Tab
} from '@mui/material';
import CheckIcon from '@mui/icons-material/Check';
import ArrowDownwardIcon from '@mui/icons-material/ArrowDownward';

function App() {
  // const currentProtocol = window.location.protocol; 
  // const currentDomain = window.location.hostname;
  // const currentPort = window.location.port;

  const [policies, setPolicies] = useState([]);
  const [selectedPolicy, setSelectedPolicy] = useState("");
  const [classifications, setClassifications] = useState([]);
  const [selectedClassification, setSelectedClassification] = useState("");
  const [categories, setCategories] = useState([]);
  const [categoriesValues, setCategoriesValues] = useState({});
  const [checkedValues, setCheckedValues] = useState({});
  const [displayType, setDisplayType] = useState(0);
  const [XMLoutput, setXMLOutput] = useState('');
  const [JSONoutput, setJSONOutput] = useState('');
  const [displayedOutput, setDisplayedOutput] = useState('')

  

  const ITEM_HEIGHT = 48;
  const ITEM_PADDING_TOP = 8;
  const MenuProps = {
    PaperProps: {
      style: {
        maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
        width: 250,
      },
    },
  };

  const handlDisplayTypeChange = (event, newValue) => {
    setDisplayType(newValue);
  };

  useEffect(() => {
    document.body.style.overflow = 'hidden';
    const fetchPolicies = async () => {
      try {
        const response = await fetch('/api/v1/policies');
        if (response.ok) {
          const policiesData = await response.json();
          setPolicies(policiesData.map((policy, index) => ({ value: policy, label: policy })));
        } else {
          console.error('Failed to fetch policies');
        }
      } catch (error) {
        console.error('Error fetching policies:', error);
      }
    };

    fetchPolicies();
  }, []);

  useEffect(() => {
    const resetValues = () => {
      setClassifications([])
      setCategories([])
      setCategoriesValues({})
      setSelectedClassification("")
      setCheckedValues({})
    }
    const fetchClassifications = async () => {
      if (selectedPolicy) {
        const classificationsResponse = await fetch(`/api/v1/classifications/${selectedPolicy}`);
          if (classificationsResponse.ok) {
            const classificationsData = await classificationsResponse.json();
            setClassifications(classificationsData.map((classification, index) => ({ value: classification, label: classification })));
        }
      }
    }
    resetValues()
    fetchClassifications()
  }, [selectedPolicy]);

  const handleSelectPolicy = async (event) => {
    setSelectedPolicy(event.target.value)    
  };
  useEffect(() => {
    const fetchCategoriesAndValues = async () => {

      if (selectedClassification) {
        const categoriesResponse = await fetch(`/api/v1/categories/${selectedPolicy}/${selectedClassification}`);
        if (categoriesResponse.ok) {
          const categoriesData = await categoriesResponse.json();
          if (categoriesData == null) {
            setCategories([])
            setCategoriesValues({})
            setCheckedValues({})
            return
          }
          else {
            setCategories(categoriesData);
          }
          const tempCategoriesValues = {};
          for (const category of categoriesData ?? []) {
            const valuesResponse = await fetch(`/api/v1/mentions/${selectedPolicy}/${selectedClassification}/${category}`);
            
            if (valuesResponse.ok) {
              const valuesData = await valuesResponse.json();
              tempCategoriesValues[category] = valuesData;
            }
          }

          setCategoriesValues(tempCategoriesValues);
          const initialCheckedValues = {};
          for (const category in tempCategoriesValues) {
              initialCheckedValues[category] = [];
          }
          setCheckedValues(initialCheckedValues);
          

        }
      }
    }
    fetchCategoriesAndValues();
  }, [selectedClassification]);

  const handleSelectClassification = async (event) => {
    setSelectedClassification(event.target.value)
  };
 
  const handleCheckboxChange = (category, values) => {
    setCheckedValues((prevCheckedValues) => {
      // Ensure we're working with a copy of the previous state
      const updatedCheckedValues = { ...prevCheckedValues };
  
      // Update the array for the category with the new values
      updatedCheckedValues[category] = values;
  
      return updatedCheckedValues;
    });
  };

  const buildJSONLabel = async() => {
    var JSONLabel = {
      "PolicyIdentifier" : selectedPolicy,
      "Classification" : selectedClassification,
      "Categories" : {}
    }
    const promises = Object.entries(checkedValues).map(async ([cat, values]) => {
      if (Array.isArray(values) && values.length > 0) {
        const typeResponse = await fetch(`/api/v1/type/${selectedPolicy}/${cat}`);
        if (typeResponse.ok) {
          const type = await typeResponse.json(); 
          JSONLabel["Categories"][cat] = {
            values: values,
            type: type
          };
        }
      }
    });
    await Promise.all(promises);
    return JSONLabel;
  }
  

  useEffect(() => {
    const updateReadableValues = async () => {
      if (selectedClassification == "") {
        setDisplayedOutput("You must first select a classification")
        return
      }
      const JSONLabel = await buildJSONLabel();
  
      // Set JSON output
      const jsonOutput = JSON.stringify(JSONLabel, null, 2);
      setJSONOutput(jsonOutput);
  
      // Fetch XML output
      fetch(`/api/v1/generate`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(JSONLabel)
      })
      .then(XMLResponse => XMLResponse.text())
      .then(xmlText => {
        // Set XML output
        setXMLOutput(xmlText);
  
        // Determine displayed output based on displayType
        switch(displayType) {
          case 0: setDisplayedOutput(xmlText); break;
          case 1: setDisplayedOutput(jsonOutput); break;
          case 2: setDisplayedOutput('To be implemented :)'); break;
        }
      });
    };
    
    updateReadableValues();
  }, [checkedValues]);

  useEffect(() => {
    const changeDisplayType  = async () => {
      if (selectedClassification == "") {
        setDisplayedOutput("You must first select a classification")
        return
      }
      switch(displayType) {
        case 0: setDisplayedOutput(XMLoutput); break
        case 1: setDisplayedOutput(JSONoutput); break;
        case 2: setDisplayedOutput('To be implemented :)')
      }
    }
    changeDisplayType()
  }, [displayType])

  return (
    <Box display="flex" flexDirection="column" height="100vh" bgcolor="#6A89A7">
  {/* Header */}
    <Box 
    display="flex"
    justifyContent="space-between"
    alignItems="center"
    padding="16px"
    bgcolor="#384959"  
    >
      <Box marginLeft="50px" >
        {/* Logo */}
        <img src="OIG5.png" alt="Logo" height="70px" />
      </Box>
      <Typography variant="h5" color="#F1F0E8" marginLeft="16px">
              SPARROW Playground
            </Typography>
      <Box marginRight="50px">
        {/* API Reference Button */}
        <Button
              variant="contained"
              component="a"
              href="/documentation/index.html"
              style={{ 
                backgroundColor: '#384959',
                border: "3px solid #E5E1DA",
                color: '#F1F0E8'

                
              }}
            >
              {'</>'} API Reference
            </Button>
      </Box>
    </Box>

  {/* Main Content */}
  <Box display="flex" height="calc(100vh - 68px)"> {/* Adjust height to account for header */}
    <Box
      width="25%"
      //border="2px solid red"
      margin="50px 50px 15px"
      display="flex"
      flexDirection="column"
      padding="16px"
      boxSizing="border-box"
      textAlign="center"
      
    > 
      <Box height="10%" display="flex" flexDirection="column" alignItems="center" justifyContent="center" margin="15px">
        <Typography  color="#F1F0E8" variant="subtitle1" sx={{ fontSize : '3ex', marginBottom: '8px' }}>
          POLICY
        </Typography>
        <FormControl required fullWidth>
          <InputLabel
            sx={{
              color: '#F1F0E8',
              '&.Mui-focused': { color: '#F1F0E8' }
            }}
          >
            Policy
          </InputLabel>
        <Select
          label="Policy"
          value={selectedPolicy}
          onChange={handleSelectPolicy}
          sx={{
            '& .MuiSelect-select': {
              fontWeight: 'bold',
              textAlign: 'center',
              color: '#F1F0E8',
            },
            '& .MuiOutlinedInput-notchedOutline': {
              borderColor: '#F1F0E8',
            },
            '&:hover .MuiOutlinedInput-notchedOutline': {
              borderColor: '#F1F0E8',
            },
            '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
              borderColor: '#F1F0E8',
            },
            '& .MuiSvgIcon-root': {
              color: '#F1F0E8',
            },
          }}
        >
          {policies.map((policy) => (
            <MenuItem key={policy.value} value={policy.value}>
              {policy.label}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
      </Box>
      <Box height="10%" display="flex" flexDirection="column" alignItems="center" justifyContent="center" margin="15px">
        <Typography color="#F1F0E8" variant="subtitle1" sx={{ fontSize : '3ex',marginBottom: '10px' }}>
          CLASSIFICATION
        </Typography>
        <FormControl fullWidth required>
          <InputLabel
          sx={{
            color: '#F1F0E8',
            '&.Mui-focused': { color: '#F1F0E8' }
          }}
          >Classification</InputLabel>
          <Select
            label="Classification"
            value={selectedClassification}
            onChange={handleSelectClassification}
            fullWidth={true}
            sx={{
              '& .MuiSelect-select': {
                fontWeight: 'bold',
                textAlign: 'center',
                color: '#F1F0E8',
              },
              '& .MuiOutlinedInput-notchedOutline': {
                borderColor: '#F1F0E8',
              },
              '&:hover .MuiOutlinedInput-notchedOutline': {
                borderColor: '#F1F0E8',
              },
              '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
                borderColor: '#F1F0E8',
              },
              '& .MuiSvgIcon-root': {
                color: '#F1F0E8',
              },
            }}
          >
            {classifications.map((classification) => (
              <MenuItem key={classification.value} value={classification.value}>
                {classification.label}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </Box>
      <Box height="70%" maxHeight="70%" display="flex" flexDirection="column" alignItems="center" justifyContent="flex-start" margin="15px">
        <Typography color="#F1F0E8" variant="subtitle1" sx={{ fontSize : '3ex', marginBottom: '10px' }}>
          CATEGORIES
        </Typography>
        <Box
          display="flex"
          flexDirection="column"
          alignItems="center"
          width="100%"
        >
        {Object.entries(categories).map(([index, category]) => (
          <FormControl sx={{ m: 1}} fullWidth>
            <InputLabel 
              id="demo-multiple-chip-label"
              sx={{
                color: '#F1F0E8',
                '&.Mui-focused': { color: '#F1F0E8' },
                bgcolor: "#6A89A7"
              }}
              
              
            >{category}</InputLabel>
            <Select
              labelId="demo-multiple-chip-label"
              id="demo-multiple-chip"dev
              fullWidth={true}
              multiple
              value={checkedValues[category] || []}
              onChange={(event) => {
                const {
                  target: { value },
                } = event;
                // Ensure value is treated as an array of selected items
                const valueArray = typeof value === 'string' ? value.split(',') : value;
                handleCheckboxChange(category, valueArray);
              }}
              sx={{
                width: '100%', // Ensure the Select component takes the full width
                '& .MuiSelect-select': {
                  fontWeight: 'bold',
                  textAlign: 'center',
                  color: '#F1F0E8',
                },
                '& .MuiOutlinedInput-notchedOutline': {
                  borderColor: '#F1F0E8',
                },
                '&:hover .MuiOutlinedInput-notchedOutline': {
                  borderColor: '#F1F0E8',
                },
                '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
                  borderColor: '#F1F0E8',
                },
                '& .MuiSvgIcon-root': {
                  color: '#F1F0E8',
                },
              }}
              input={<OutlinedInput id="select-multiple-chip" label="Chip" />}
              renderValue={(selected) => (
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {selected.map((value) => (
                    <Chip key={value} label={value} sx={{
                      color: 'white', 
                    }}/>
                  ))}
                </Box>
              )}
              MenuProps={MenuProps}
            >
              {categoriesValues[category]?.map((value) => (
                <MenuItem key={value} value={value}>
                  {value}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        ))}
        </Box>
      </Box>
      
    </Box>

    {/* Right Zone */}
    <Box width="75%" padding="16px" boxSizing="border-box" margin="50px 50px 50px 50px">
      <Tabs value={displayType} onChange={handlDisplayTypeChange} centered sx={{
        '& .MuiTab-root': {
          color: '#F1F0E8 !important',
          fontSize: '20px', // Adjust the font size as needed
        },
        '& .Mui-selected': {
          color: '#F1F0E8 !important',
          fontWeight: 'bold',
          fontSize: '20px', // Adjust the font size as needed
        },
        '& .MuiTabs-indicator': {
          backgroundColor: '#F1F0E8',
        },
      }}>
        <Tab label="XML" />
        <Tab label="JSON" />
        <Tab label="SVG" />
      </Tabs>
      <Box
        sx={{
          margin : '10px',
          overflow: 'auto',
          height: '90%',
          flexDirection: 'column',
          border: '1px solid rgba(0, 0, 0, 0.23)', // Outline the Box
          borderRadius: '4px', // Optional: for rounded corners
          position: 'relative', // For positioning the arrow
        }}
      >
        <TextField
          multiline
          fullWidth
          //variant="outlined"
          value={displayedOutput}
          InputProps={{
            style: {
              fontFamily: 'monospace',
              padding: '10px',
              color: '#F1F0E8',
            },
            disableUnderline: true,
            readOnly: true,
          }}
          sx={{
            flexGrow: 1,
            overflow: 'auto',
            '& .MuiOutlinedInput-notchedOutline': {
              border: 'none',
            },
          }}
        />
        
      </Box>
    </Box>
  </Box>
</Box>

  );
}

export default App;
