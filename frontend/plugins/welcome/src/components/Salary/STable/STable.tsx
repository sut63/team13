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
import { DefaultApi } from '../../../api/apis'; // Api Gennerate From Command
import { EntAssessment } from '../../../api/models/EntAssessment';
import { EntEmployee } from '../../../api/models/EntEmployee';
import { EntPosition } from '../../../api/models/EntPosition';
import { EntSalary } from '../../../api/models/EntSalary';
//import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
//import { ContentHeader } from '@backstage/core';
import { createMuiTheme } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
import Swal from 'sweetalert2';
import ComponanceTable from '../SalaryTable';


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

    primary: {
      light: '#757ce8',
      main: '#3f50b5',
      dark: '#002884',
      contrastText: '#fff',
    },
  }),
);

const theme = createMuiTheme({
  palette: {
    primary: {
      main: purple[500],
    },
    secondary: {
      main: '#f44336',
    },
  },
});






export default function MenuAppBar() {

    const classes = useStyles();
  




 function HomeIcon(props:any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }

  return (
    <div className={classes.root}>
      
      <AppBar
        component="div"
        color= 'secondary'
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              <Typography color="inherit" variant="h3" component="h2">
                ระบบบันทึกเงินเดือนพนักงาน
              </Typography>
            </Grid>
            <Grid item>
                <IconButton 
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/SplitsystemManager"
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
            
            
          </Grid>
          
        </Toolbar>
        
        <ComponanceTable></ComponanceTable>
        
      </AppBar>
      
            
       <Grid item xs={2}> </Grid>
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="inherit"
        position="static"
        elevation={1}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={5}>

          <Link component={RouterLink} to="/Salary">
                <Button variant="contained" color="secondary" style={{ marginLeft : 1350}}>
                Back
            </Button>
            </Link>
            
        </Grid>
        </Toolbar>
      </AppBar>
      
    
    
    </div>
  );
 }