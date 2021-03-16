import React, { useState, useEffect } from 'react';
import { createMuiTheme, createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../../api/apis';
import { EntEmployeeWorkingHours } from '../../../api/models/EntEmployeeWorkingHours';
import moment from "moment";
import { 
  SvgIcon, 
  AppBar, 
  Toolbar, 
  Grid, 
  Typography, 
  IconButton, 
  Link 
} from '@material-ui/core';
import { purple } from '@material-ui/core/colors';
import { Link as RouterLink } from 'react-router-dom';

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
    table: {
      minWidth: 600,
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


export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [EmployeeWorkingHourss, setEmployeeWorkingHourss] = useState<EntEmployeeWorkingHours[]>(Array);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listEmployeeworkinghours({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployeeWorkingHourss(res);
    };
    getUsers();
  }, [loading]);

  const deleteEmployeeWorkingHours = async (id: number) => {
    const res = await api.deleteEmployeeworkinghours({ id: id });
    setLoading(true);
  };

  function HomeIcon(props:any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }
  console.log(EmployeeWorkingHourss)
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
        
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">No.</TableCell>
            <TableCell align="center">ชื่อพนักงาน</TableCell>
            <TableCell align="center">เลขบัตรประชาชน</TableCell>
            <TableCell align="center">รหัสพนักงาน</TableCell>
            <TableCell align="center">วันที่เข้าทำงาน</TableCell>
            <TableCell align="center">เวลาเริ่มงาน</TableCell>
            <TableCell align="center">เวลาเลิกงาน</TableCell>
            <TableCell align="center">หน้าที่ที่รับผิดชอบ</TableCell>
            <TableCell align="center">ค่าจ้าง</TableCell>
            <TableCell align="center">Manage</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {EmployeeWorkingHourss === undefined
            ? null
            : EmployeeWorkingHourss.map((item:any) => (
              <TableRow key={item.id}>
                <TableCell align="center">{item.id}</TableCell>
                <TableCell align="center">{item.edges?.employee?.name}</TableCell>
                <TableCell align="center">{item.iDNumber}</TableCell>
                <TableCell align="center">{item.codeWork}</TableCell>
                <TableCell align="center">{item.edges?.day?.day}</TableCell>
                <TableCell align="center">{moment(item.edges?.startwork?.startWork).format("LT")}</TableCell>
                <TableCell align="center">{moment(item.edges?.endwork?.endWork).format("LT")}</TableCell>
                <TableCell align="center">{item.edges?.role?.role}</TableCell>
                <TableCell align="center">{item.wages}</TableCell>
                <TableCell align="center">
                  <Button
                    onClick={() => {
                      deleteEmployeeWorkingHours(item.id);
                    }}
                    style={{ marginLeft: 10 }}
                    variant="contained"
                    color="secondary"
                  >
                    Delete
               </Button>
                </TableCell>
              </TableRow>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
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

          <Link component={RouterLink} to="/EmployeeWorkingHours">
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
