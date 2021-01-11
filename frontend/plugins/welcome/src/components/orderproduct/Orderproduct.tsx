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
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
//import { EntPhysician } from '../../api/models/EntPhysician';
//import { EntMedicalType } from '../../api/models/EntMedicalType';
//import { EntMedicalEquipment } from '../../api/models/EntMedicalEquipment';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
//import { ContentHeader } from '@backstage/core';
import { Alert } from '@material-ui/lab';

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
  }),
);



function Copyright() {
  return (
    <Typography variant="body2" color="inherit" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Poommin Website
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
  
  //const [physicians, setPhysicians] = useState<EntPhysician[]>([]);
  //const [medicalEquipments, setMedicalEquipments] = useState<EntMedicalEquipment[]>([]);
  //const [medicalTypes, setMedicalTypes] = useState<EntMedicalType[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);


  const [physicianid, setPhysicianid] = useState(Number);
  const [medicalEquipmentid, setMedicalEquipmentid] = useState(Number);
  const [medicalTypeid, setMedicalTypeid] = useState(Number);
  const [medicalEquipmentstockid, setMedicalEquipmentstockid] = useState(Number);
  const [datetime, setDatetime] = useState(String);

 let stock = Number(medicalEquipmentstockid)
 let nameEquipmentID = Number(medicalEquipmentid)
 let typeEquipmentID =Number(medicalTypeid)
 let physicianID = Number(physicianid)

 console.log(physicianID)
  useEffect(() => {

    const getmedicalEquipments = async () => {
 
     // const pa = await api.listMedicalequipment({ limit: 10, offset: 0 });
      setLoading(false);
     // setMedicalEquipments(pa);
    };
    getmedicalEquipments();
 
    const getPhysicians = async () => {
 
   // const us = await api.listPhysician({ limit: 10, offset: 0 });
      setLoading(false);
     // setPhysicians(us);
    };
    getPhysicians();
 
    const getmedicalTypes = async () => {
 
     //const sp = await api.listMedicaltype({ limit: 10, offset: 0 });
       setLoading(false);
     //  setMedicalTypes(sp);
     };
     getmedicalTypes();

     
  }, [loading]);
  
const systemequipment = {
                 
  physicianID  , 
  nameEquipmentID ,   
  stock , 
  typeEquipmentID ,
  addedtime :datetime   + ":00+07:00"
}
console.log(systemequipment)
const createSystemequipment = async () => {
 
 console.log(medicalEquipmentid)
//const res:any = await api.createSystemequipment({ systemequipment : systemequipment});
setStatus(true);
if (res.id != ''){
 setAlert(true);
} else {
 setAlert(false);
}

const timer = setTimeout(() => {
 setStatus(false);
}, 1000);
};
  
const physician_id_handleChange = (event: any)=> {
  setPhysicianid(event.target.value);
   }; 

  const Equipment_id_handleChange = (event:any) => {
    setMedicalEquipmentid(event.target.value);
   };

  const Type_id_handleChange = (event: any) => {
    setMedicalTypeid(event.target.value);
  }
  const Stock_id_handleChange = (event: any) => {
    setMedicalEquipmentstockid(event.target.value);
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
            <Typography color="inherit"  component="h1">
                ระบบสั่งสินค้าเข้ามาในคลัง
              </Typography>
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
                ชื่อแพทย์
              </Typography>
            </Grid>
            <Grid item xs={2}>
            
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อเครื่องมือแพทย์
              </Typography>
            </Grid>
            <Grid item xs={2}>
              
            
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ประเภทของอุปกรณ์
              </Typography>
            </Grid>
            <Grid item xs={2}>
            
            



            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                จำนวน/ชิ้น
              </Typography>
            </Grid>
            <Grid item xs={2}>
            <Paper >
                <TextField id="outlined-number" type='number'  InputLabelProps={{
                  shrink: true,}}label="กรุณาใส่จำนวน" variant="outlined"
                  onChange = {Stock_id_handleChange}
                  />
                  </Paper>
                  
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
                  createSystemequipment();
                }}
                
                startIcon={<SaveIcon 
                />}
              >
                Save
              </Button>
            </Grid>
            <Grid item xs={2}></Grid>
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