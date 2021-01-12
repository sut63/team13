import React, { useState,Component, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import Hidden from '@material-ui/core/Hidden';
import { Link as RouterLink } from 'react-router-dom';
import Link from '@material-ui/core/Link';
import Avatar from '@material-ui/core/Avatar';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import SvgIcon from '@material-ui/core/SvgIcon';
import SaveIcon from '@material-ui/icons/Save';
import Select from '@material-ui/core/Select';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntAssessment } from '../../api/models/EntAssessment';
import { EntEmployee } from '../../api/models/EntEmployee';
import { EntPosition } from '../../api/models/EntPosition';
import { EntSalary } from '../../api/models/EntSalary';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
//import { ContentHeader } from '@backstage/core';
import { Alert } from '@material-ui/lab';
import {
    FormControl,
    InputLabel,
    Box,
  } from '@material-ui/core';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Menu from '@material-ui/core/Menu';


const lightColor = 'rgba(255, 255, 255, 0.7)';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    iconButtonAvatar: {
      padding: 4,
    },
    secondaryBar: {
      zIndex: 0,
    },
    button: {
      borderColor: lightColor,
    },
    title: {
      flexGrow: 1,
    },
    textField: {
      width: 200,
    },
    margin: {
        marginRight: theme.spacing(2),
      },
    
  }),
);



function Copyright() {
  return (
    <Typography variant="body2" color="inherit" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Nontakorn Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

export default function MenuAppBar() {
  
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const classes = useStyles();
  const profile = { givenName: 'to Software Analysis 63' };
  const api = new DefaultApi();
  
  const [assessments, setAssessments] = useState<EntAssessment[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [positions, setPositions] = useState<EntPosition[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);


  const [assessmentid, setAssessmentid] = useState(Number);
  const [employeeid, setEmployeeid] = useState(Number);
  const [positionid, setPositionid] = useState(Number);
  const [salaryNumber, setSalaryNumber] = useState(Number);
  const [bonusNumber, setBonus] = useState(Number);
  const [datetime, setDatetime] = useState(String);

 let assessmentID = Number(assessmentid) 
 let employeeID  = Number(employeeid)
 let positionID =Number(positionid)
 let salarys  = Number(salaryNumber)
 let bonus  = Number(bonusNumber)

 console.log(employeeID)
 useEffect(() => {

  const getEmployees = async () => {

    const em = await api.listEmployee({ limit: 10, offset: 0 });
    setLoading(false);
    setEmployees(em);
  };
  getEmployees();

  const getAssessments = async () => {

  const tp = await api.listAssessment({ limit: 10, offset: 0 });
    setLoading(false);
    setAssessments(tp);
  };
  getAssessments();

  const getPositions = async () => {

   const pr = await api.listPosition({ limit: 10, offset: 0 });
     setLoading(false);
     setPositions(pr);
   };
   getPositions();

   
}, [loading]);
  
const Salary = {
                 
    assessmentID , 
    bonus ,   
    employeeID , 
    positionID ,
    salarys ,
    salaryDate :datetime   + ":00+07:00"
}
console.log(Salary)
const createSalary = async () => {
 
//console.log()
const res:any = await api.createSalary({ salary : Salary });
setStatus(true);
if (res.id != ''){
    setAlert(true);
    window.location.reload(false);
} else {
    setAlert(false);
    setStatus(true);
}

const timer = setTimeout(() => {
 setStatus(false);
}, 1000);
};
  
const employee_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployeeid(event.target.value as number);
  };

  const assessment_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAssessmentid(event.target.value as number);
  };

  const position_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPositionid(event.target.value as number);
  };
  const handleSalaryChange = (event: any) => {
    setSalaryNumber(event.target.value as number);
  };

  const handleBonusChange = (event: any) => {
    setBonus(event.target.value as number);
  };

  const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  }; 
  

 function HomeIcon(props:any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }

  return (
    <div className={classes.root}>
      <AppBar color="primary" position="sticky" elevation={0}>
        <Toolbar>
          <Grid container spacing={1} alignItems="center">
            <Hidden smUp>
            <Grid item>
                
              </Grid>
            </Hidden>
            <Grid item xs />
            <Grid item>
            
            </Grid>
            
            <Grid item>
                <IconButton 
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/afterlogin"
                >    
                <HomeIcon color="inherit" />
                </IconButton>
                </Grid>
                <Grid item>
            <Button className={classes.button} variant="outlined" color="inherit" 
            size="small" component={RouterLink}
            to="/">
                logout
              </Button>
                </Grid>  
            <Grid item>
            
            </Grid>
            <Grid item>
              <IconButton color="inherit" className={classes.iconButtonAvatar}>
                <Avatar src="o" alt="P" />
              </IconButton>
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
      
      <AppBar
        component="div"
        color="primary"
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              <Typography color="inherit" variant="h2" component="h2">
                ระบบบันทึกเงินเดือนพนักงงาน
              </Typography>
            </Grid>
            <Grid item>
            
                </Grid>  
            
            <Grid item>
            
            </Grid>
            
          </Grid>
        </Toolbar>
      </AppBar>
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="primary"
        position="static"
        elevation={0}
      >
        <Tabs value={0} textColor="inherit">
          <Tab textColor="inherit" label="ADD Data" />     
        </Tabs>
        
      </AppBar>
            
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
       
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="inherit"
        position="static"
        elevation={1}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={4}>
            <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>
        
            
            
            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อพนักงาน
              </Typography>
            </Grid>
            <Grid item xs={2}>
              
            <Select
               labelId="employee_id-label"
               label="employee"
               id="employee_id"
               onChange={employee_id_handleChange}
               style = {{width: 200}}
               >
               {employees.map((item:EntEmployee)=>
               <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>)}
             </Select>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ตำแหน่ง
              </Typography>
            </Grid>
            <Grid item xs={2}>
              
            <Select
               labelId="position_id-label"
               label="position"
               id="position_id"
               onChange={position_id_handleChange}
               style = {{width: 200}}
               >
               {positions.map((item:EntPosition)=>
               <MenuItem key={item.id} value={item.id}>{item.position}</MenuItem>)}
             </Select>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ผลการประเมิน
              </Typography>
            </Grid>
            <Grid item xs={2}>
            
            <Select
               labelId="assessment_id-label"
               label="assessment"
               id="assessment_id"
               onChange={assessment_id_handleChange}
               style = {{width: 200}}
               >
               {assessments.map((item:EntAssessment)=>
               <MenuItem key={item.id} value={item.id}>{item.assessment}</MenuItem>)}
             </Select>

             

            </Grid>

            
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                เงินเดือน
              </Typography>
            </Grid>
            <Grid item xs={2}>
              
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"

            ><TextField
                id="salary"
                label="salary"
                variant="outlined"
                type="string"
                size="medium"
                value={salaryNumber}
                onChange={handleSalaryChange }
                style={{  marginRight :300,width: 300 }}
              />
              </FormControl>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>


            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                โบนัส
              </Typography>
            </Grid>
            <Grid item xs={2}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"

            ><TextField
                id="price"
                label="Price"
                variant="outlined"
                type="string"
                size="medium"
                value={bonusNumber}
                onChange={handleBonusChange }
                style={{  marginRight :300,width: 300 }}
              />
            </FormControl>
                  
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                วันที่/เวลาที่บันทึก
              </Typography>
            </Grid>
            <Grid item xs={2}>
               <form  noValidate>
               <TextField
                      id="date"
                      label="DateTime"
                      type="datetime-local"
                      value={datetime}
                      onChange={handleDatetimeChange}
                      //defaultValue="2017-05-24"
                      className={classes.textField}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
              </form>
              
                  
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Button
              
                variant="contained"
                color="primary"
                size="large"
                className={classes.button}
                onClick={() => {
                  createSalary();
                }}
                
                startIcon={<SaveIcon 
                />}
              >
                Save
              </Button>
            </Grid>
            
            <Grid item xs={2}>
            <Link component={RouterLink} to="/SalaryTable">
                <Button variant="contained" color="secondary" style={{ marginLeft : 40}}>
                หน้าหลัก
            </Button>
            </Link>

            </Grid>
            <Grid item xs={2}> </Grid>
            <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>
            
        </Grid>
        </Toolbar>
      </AppBar>
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="primary"
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>  
            <Copyright />
              
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
    
    </div>
  );
 }