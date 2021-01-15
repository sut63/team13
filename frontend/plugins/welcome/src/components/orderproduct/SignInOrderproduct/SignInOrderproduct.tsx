import React, { FC , useEffect, useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { Cookies } from './Cookie';
import { DefaultApi } from '../../../api/apis'; 
import { EntManager } from '../../../api';
import { Link as RouterLink } from 'react-router-dom';

import { Alert } from '@material-ui/lab'; // alert
import { Content, ContentHeader } from '@backstage/core';



const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignInManager: FC<{}> = () => {

  const classes = useStyles();
  const api = new DefaultApi();
  var ck = new Cookies();
  var check : boolean
  const [path, setPath] = React.useState("");
  const [alert, setAlert] = useState(Boolean);
  const [status, setStatus] = useState(false);
  
  // list ManagerStaff
  const [manager,setManager] = React.useState<EntManager[]>([])
  const listManager = async() => {
        const res = await api.listManager({})
        setManager(res)
  }

  // setEmail
  const [email, setEmail] = React.useState()
  const handleEmail = (event : any) => {
      setEmail(event.target.value)
  }
/*
  const [name, setName] = React.useState()
  const handleName = (event : any) => {
      setName(event.target.value)
  }*/

  // setPassword
  const [password, setPassword] = React.useState()
  const handlePassword = (event : any) => {
      setPassword(event.target.value)
  }

  // handleCookies
  function handleCookies() {
    check = ck.CheckLogin(manager,email,password)
    console.log("check => "+check)
    if(check === true){
      setAlert(true);
      history.pushState('', '', '/SplitsystemManager');
      ck.SetCookie("email",email,30)
      ck.SetCookie("id",ck.SetID(manager,email,password),30)
      ck.SetCookie("name",ck.SetName(manager,email,password),30)
      
      window.location.reload(false)
    }else if(check === false){
      setAlert(false);
      //alert("The wrong password or email was entered.!!!")
      //setPath("/")
    }
    setStatus(true);
    const timer = setTimeout(() => {
      setStatus(false);
    }, 3000);
  };
  // useEffect
  useEffect(() => {
      listManager()
  },[])

  return (
    <Content>
      <ContentHeader title="Login Manager">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">เข้าสู่ระบบสำเร็จ</Alert>
              ) : (
                  <Alert severity="error" >
                    ไม่พบข้อมูลในระบบ
                  </Alert>
                )}
            </div>
          ) : null}
      </ContentHeader>
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={handleEmail}
          />
         
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            onChange={handlePassword}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleCookies}
            component={RouterLink}
            to={path}
          >
            Sign In
          </Button>

          <Button
            type="submit"
            fullWidth
            variant="contained"
            className={classes.submit}
            component={RouterLink}
                  to="/"
            
          >
            Back
          </Button>
          
        </form>
      </div>
    </Container>
    </Content>
  );
};

export default SignInManager;
