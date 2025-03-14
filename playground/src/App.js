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
  Chip
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
  /*
  const handleCategorySelectionChange = (event) => {
    const {
      target: { value },
    } = event;
    setCheckedValues(
      // On autofill we get a stringified value.

      typeof value === 'string' ? value.split(',') : value,
    );
  };*/
  const handleCheckboxChangee = (category, value) => {
    setCheckedValues((prevCheckedValues) => {
      // Ensure we're working with a copy of the previous state
      const updatedCheckedValues = { ...prevCheckedValues };
  
      // Initialize the array for the category if it doesn't exist
      if (!updatedCheckedValues[category]) {
        updatedCheckedValues[category] = [];
      }
  
      // Toggle the value in the array
      if (updatedCheckedValues[category].includes(value)) {
        updatedCheckedValues[category] = updatedCheckedValues[category].filter((item) => item !== value);
      } else {
        updatedCheckedValues[category] = [...updatedCheckedValues[category], value];
      }
  
      return updatedCheckedValues;
    });
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
    const handleLogCheckedValues = () => {
      console.log('Checked Values:', checkedValues);
    };

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
      <Box width="75%" border="2px solid blue" padding="16px" boxSizing="border-box">
      
      </Box>
    </Box>
  );
}

export default App;
