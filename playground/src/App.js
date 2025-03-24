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
import PlayArrowIcon from '@mui/icons-material/PlayArrow';

function App() {
  const [sparrowUrl, setSparrowUrl] = useState("");
  const [sparrowUrlValid, setSparrowUrlValid] = useState(null);
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

  /*const handleOutputChange = (event) => {
    setOutput(event.target.value);
  };*/

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

  const handleTestClick = async () => {
    try {
      const response = await fetch(sparrowUrl + '/api/v1/policies');
      if (response.ok) {
        setSparrowUrlValid(true);
      } else {
        setSparrowUrlValid(false);
        console.error('Error:', response.statusText);
      }
    } catch (error) {
      setSparrowUrlValid(false);
      console.error('Error:', error);
    }
  };

  const handleValidate = async () => {
    if (sparrowUrlValid) {
      const policiesResponse = await fetch(`${sparrowUrl}/api/v1/policies`);
        if (policiesResponse.ok) {
          const policiesData = await policiesResponse.json();
          setPolicies(policiesData.map((policy, index) => ({ value: policy, label: policy })));
      }
    }
  };

  useEffect(() => {
    const fetchClassifications = async () => {
      if (selectedPolicy) {
        const classificationsResponse = await fetch(`${sparrowUrl}/api/v1/classifications/${selectedPolicy}`);
          if (classificationsResponse.ok) {
            const classificationsData = await classificationsResponse.json();
            setClassifications(classificationsData.map((classification, index) => ({ value: classification, label: classification })));
        }
      }
    }
    fetchClassifications()
  }, [selectedPolicy]);

  const handleSelectPolicy = async (event) => {
    if (sparrowUrlValid) {
      setSelectedPolicy(event.target.value)
    }     
  };
  useEffect(() => {
    const fetchCategoriesAndValues = async () => {
      if (selectedClassification) {
        const categoriesResponse = await fetch(`${sparrowUrl}/api/v1/categories/${selectedPolicy}/${selectedClassification}`);
        if (categoriesResponse.ok) {
          const categoriesData = await categoriesResponse.json();
          setCategories(categoriesData);
          const tempCategoriesValues = {};
          for (const category of categoriesData) {
            const valuesResponse = await fetch(`${sparrowUrl}/api/v1/mentions/${selectedPolicy}/${selectedClassification}/${category}`);
            
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
  useEffect(() => {
    console.log(categories);
    console.log(categoriesValues);
    console.log(checkedValues);
  }, [categoriesValues, checkedValues]);

  const handleSelectClassification = async (event) => {
    if (sparrowUrlValid) {
      setSelectedClassification(event.target.value)
    }     
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
        const typeResponse = await fetch(`${sparrowUrl}/api/v1/type/${selectedPolicy}/${cat}`);
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
  const handleLogCheckedValues = async () => {
    /*
    if (displayType === 0) {
   
      const JSONLabel = await buildJSONLabel();
      fetch(`${sparrowUrl}/api/v1/generate`, {
        method: 'POST', // HTTP method
        headers: {
            'Content-Type': 'application/json' // Specify JSON format
        },
        body:JSON.stringify(JSONLabel)
    }).then(XMLResponse => XMLResponse.text()) // Convert XML response to text
      .then(xmlText => setOutput(xmlText)) // Set resolved text
    }

    if (displayType === 1) {
   
      const JSONLabel = await buildJSONLabel();
      
      setOutput(JSON.stringify(JSONLabel, null, 2))
    }
    else {
      setOutput('To be implemented :)')
    }*/
  };

  useEffect(() => {
    const updateReadableValues = async () => {
      const JSONLabel = await buildJSONLabel();
  
      // Set JSON output
      const jsonOutput = JSON.stringify(JSONLabel, null, 2);
      setJSONOutput(jsonOutput);
  
      // Fetch XML output
      fetch(`${sparrowUrl}/api/v1/generate`, {
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
      
      switch(displayType) {
        case 0: setDisplayedOutput(XMLoutput); break
        case 1: setDisplayedOutput(JSONoutput); break;
        case 2: setDisplayedOutput('To be implemented :)')
      }
    }
    changeDisplayType()
  }, [displayType])

  return (
    <Box display="flex" height="100vh">
      <Box
        width="25%"
        border="2px solid red"
        margin= "50px 50px 15px"
        display="flex"
        flexDirection="column"
        padding="16px"
        boxSizing="border-box"
      >
        <Box height="10%" display="flex" alignItems="center" justifyContent="space-between">
          <TextField 
            label="URL" 
            value={sparrowUrl}
            onChange={(e) => setSparrowUrl(e.target.value)}
            sx={{ 
              flexGrow: 1, 
              marginRight: '8px',
              '& .MuiOutlinedInput-root': {
                '& fieldset': {
                  borderColor: sparrowUrlValid === null ? 'inherit' : sparrowUrlValid ? 'green' : 'red',
                },
                '&:hover fieldset': {
                  borderColor: sparrowUrlValid === null ? 'inherit' : sparrowUrlValid ? 'green' : 'red',
                },
                '&.Mui-focused fieldset': {
                  borderColor: sparrowUrlValid === null ? 'inherit' : sparrowUrlValid ? 'green' : 'red',
                },
              },
            }}
          />
          <Button variant="contained" color="primary" sx={{ marginRight: '8px' }} onClick={handleTestClick}>
            <PlayArrowIcon />
          </Button>
          <Button variant="contained" color="success" onClick={handleValidate}>
            <CheckIcon />
          </Button>
        </Box>
        <Box height="10%" display="flex" flexDirection="column" alignItems="center" justifyContent="center" margin="15px">
          <Typography variant="subtitle1" sx={{ marginBottom: '8px' }}>
            POLICY
          </Typography>
          <FormControl required fullWidth>
            <InputLabel>Policy</InputLabel>
            <Select 
              label="Policy"
              value={selectedPolicy}
              onChange={handleSelectPolicy}
              sx={{
                '& .MuiSelect-select': {
                  fontWeight: 'bold',
                  textAlign: 'center',
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
          <Typography variant="subtitle1" sx={{ marginBottom: '10px' }}>
            CLASSIFICATION
          </Typography>
          <FormControl fullWidth>
            <InputLabel>Classification</InputLabel>
            <Select 
              label="Classification"
              value={selectedClassification}
              onChange={handleSelectClassification}
            >
              {classifications.map((classification) => (
                <MenuItem key={classification.value} value={classification.value}>
                  {classification.label}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Box>
        <Box height="70%" display="flex" flexDirection="column" alignItems="center" justifyContent="center" margin="15px">
        <Typography variant="subtitle1" sx={{ marginBottom: '10px' }}>
          CATEGORIES
        </Typography>
        {Object.entries(categories).map(([index, category]) => (
          <FormControl sx={{ m: 1, width: 300 }}>
            <InputLabel id="demo-multiple-chip-label">{category}</InputLabel>
            <Select
             labelId="demo-multiple-chip-label"
             id="demo-multiple-chip"
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
             input={<OutlinedInput id="select-multiple-chip" label="Chip" />}
             renderValue={(selected) => (
               <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                 {selected.map((value) => (
                   <Chip key={value} label={value} />
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
        <Box display="flex" justifyContent="center" marginTop="15px">
          <Button variant="contained" color="secondary" onClick={handleLogCheckedValues}>
            Log Checked Values
          </Button>
        </Box>
        
      </Box>

      {/* Zone de droite */}
      <Box width="75%" border="2px solid blue" padding="16px" boxSizing="border-box" margin= "50px 50px 50px 50px">
      <Tabs value={displayType} onChange={handlDisplayTypeChange} centered>
        <Tab label="XML" />
        <Tab label="JSON" />
        <Tab label="SVG"/>
      </Tabs>
      <Box
      sx={{
        border: '1px solid #ccc',
        borderRadius: '4px',
        overflow: 'auto',
        height: '100%', // Ensure the container takes full height
        
      }}
    >
      <TextField
        multiline
        fullWidth
        variant="outlined"
        value={displayedOutput}
        //onChange={handleOutputChange}
        InputProps={{
          style: {
            fontFamily: 'monospace',
            padding: '10px',
          },
          disableUnderline: true,
          readOnly: true, // Make the text area read-only
        }}
      />
    </Box>
      </Box>
    </Box>
  );
}

export default App;
