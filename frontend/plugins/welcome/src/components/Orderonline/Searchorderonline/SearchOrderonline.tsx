import React, { useState } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import { Link as RouterLink } from 'react-router-dom';
import Avatar from '@material-ui/core/Avatar';
import { DefaultApi } from '../../../api/apis';
import Swal from 'sweetalert2';
import { Cookiesonline } from '../SignInOrderonline/Cookie'
import SearchIcon from '@material-ui/icons/Search';
import { EntOrderonline } from '../../../api';;
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import moment from 'moment';
import { TextField } from '@material-ui/core';

const lightColor = 'rgba(255, 255, 255, 0.7)';
const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

const useStyles = makeStyles((theme: Theme) =>

  createStyles({
    table: {
      minWidth: 650,
    },
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

export default function MenuAppBar() {

  var ck = new Cookiesonline()
  var cookieEmail = ck.GetCookie()
  var cookieID = ck.GetID()
  var cookieName = ck.GetName()
  const classes = useStyles();
  const api = new DefaultApi();

  const [productid, setProductid] = useState(String);

  let customerID = Number(cookieID)
  console.log(customerID)

  const [orderonlines, setOrderonlines] = useState<EntOrderonline[]>();

  const orderonline = {
    customerID
  }
  console.log(orderonline)

  const Product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProductid(event.target.value as string);
  };

  var lenOrderonline: number
  
  const getCheckinsorder = async () => {
    const res = await api.searchOrderonline({ userid: customerID, name: productid })
    setOrderonlines(res)
    console.log(res)
    lenOrderonline = res.length
    if (lenOrderonline > 0) {
      //setOpen(true)
      Toast.fire({
        icon: 'success',
        title: 'ค้นหาข้อมูลสำเร็จ',
      })
    } else {
      //setFail(true)
      Toast.fire({
        icon: 'error',
        title: 'ค้นหาข้อมูลไม่พบ',
      })
    }
  }

  return (

    <Page theme={pageTheme.home}>
      <Header
        title=" Welcome to Orderonline "
        subtitle="Select Product you want to search."
      >

        <TableCell align="right">
          <Button
            style={{ marginLeft: 1 }}
            component={RouterLink}
            to="/Orderonline"
            variant="contained"
          >
            Back
             </Button>
        </TableCell>    

        <IconButton color="inherit" className={classes.iconButtonAvatar}>
          <Avatar src='o' alt={cookieEmail} />
        </IconButton>

      </Header>
      <Content>
        <div className={classes.root}>
          <form noValidate autoComplete="off">


            <TableCell align="left"></TableCell>

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
                  <Grid item xs={12}>
                  </Grid>
                  <Grid item xs={2}></Grid>
                  <Grid item xs={2}></Grid>
                  <Grid item xs={2}>
                    <Typography color="primary" variant="h6" component="h1">
                      ชื่อผู้ใช้ระบบค้นหา
                    </Typography>
                  </Grid>      
                  <Grid item xs={2}>
                    <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
                  </Grid>
                  <Grid item xs={12}></Grid>
                  <Grid item xs={2}></Grid>

                  <Grid item xs={2}></Grid>
                  
                  <Grid item xs={2}>
                  <Typography color="primary" variant="h6" component="h1">
                     ชื่อสินค้าที่ต้องการค้นหา
                  </Typography>
                  </Grid>
                  <Grid item xs={2}>

                  <TextField
                  id="product"
                  label="product"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={productid}
                  onChange={Product_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                />
            </Grid>

                  <TableCell align="center">
                    <Button
                      style={{ marginLeft: 500 }}
                      variant="contained"
                      color="primary"
                      size="large"
                      className={classes.button}
                      onClick={() => {
                        getCheckinsorder();
                      }}

                      startIcon={<SearchIcon
                      />}
                    >
                      Search
              </Button>
                  </TableCell>
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
            </AppBar>
            <TableContainer component={Paper}>
              <Table className={classes.table} aria-label="simple table">
                <TableHead>
                  <TableRow>
                    
                    <TableCell align="center">Customer</TableCell>
                    <TableCell align="center">Typeproduct</TableCell>
                    <TableCell align="center">Product</TableCell>
                    <TableCell align="center">Stock</TableCell>
                    <TableCell align="center">Paymentchannel</TableCell>
                    <TableCell align="center">Accountnumber</TableCell>
                    <TableCell align="center">CVV</TableCell>
                    <TableCell align="center">Time</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {orderonlines === undefined
                    ? null
                    : orderonlines.map((item: any) => (
                      <TableRow key={item.productid}>
                     
                        <TableCell align="center">{item.edges?.customer?.name}</TableCell>
                        <TableCell align="center">{item.edges?.typeproduct?.typeproduct}</TableCell>
                        <TableCell align="center">{item.edges?.product?.nameProduct}</TableCell>
                        <TableCell align="center">{item.stock}</TableCell>
                        <TableCell align="center">{item.edges?.paymentchannel?.bank}</TableCell>
                        <TableCell align="center">{item.accountnumber}</TableCell>
                        <TableCell align="center">{item.cvv}</TableCell>
                        <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH:mm')}</TableCell>
                      </TableRow>
                    ))}
                </TableBody>
              </Table>
            </TableContainer>
          </form>
        </div>
      </Content>
    </Page >
  );
}

