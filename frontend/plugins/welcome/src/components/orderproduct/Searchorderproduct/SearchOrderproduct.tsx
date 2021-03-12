import React, { useState, useEffect } from 'react';
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
import { DefaultApi } from '../../../api/apis';
import Select from '@material-ui/core/Select';
import { EntProduct } from '../../../api/models/EntProduct';
import Swal from 'sweetalert2';
import { Cookies } from '../SignInOrderproduct/Cookie'
import SearchIcon from '@material-ui/icons/Search';
import { EntOrderproduct } from '../../../api';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import { FormControl, InputLabel, TextField } from '@material-ui/core';

const lightColor = 'rgba(255, 255, 255, 0.7)';
const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  width: '400px',
  padding: '100px',
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
//testupgithub

export default function MenuAppBar() {

  var ck = new Cookies()
  var cookieEmail = ck.GetCookie()
  var cookieID = ck.GetID()
  var cookieName = ck.GetName()
  const classes = useStyles();
  const api = new DefaultApi();
  const [search, setSearch] = useState(false);
  const [checkProductName, setProductNames] = useState(false);
  const [products, setProducts] = useState<EntProduct[]>([]);
  const [loading, setLoading] = useState(true);
  const [productid, setProductid] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }



  const [orderproducts, setOrderproducts] = useState<EntOrderproduct[]>([]);


  const deleteSystemequipments = async (id: number) => {
    const res = await api.deleteOrderproduct({ id: id });
    setLoading(true);
  };
  useEffect(() => {
    const getproducts = async () => {
      const pr = await api.listProduct({ offset: 0 });
      setLoading(false);
      setProducts(pr);
    };
    getproducts();
    const getOrderproduct = async () => {
      const sp = await api.listOrderproduct({ offset: 0 });
      setLoading(false);
      setOrderproducts(sp);
    };
    getOrderproduct();
  }, [loading]);


  const Product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProductid(event.target.value as string);
    setProductNames(false);
    setSearch(false);
  }

  const getCheckOrderproduct = async () => {
    var check = false;
    orderproducts.map(item => {
      if (productid != "") {
        if (item.edges?.product?.nameProduct?.startsWith(productid)) {
          setProductNames(true);
          alertMessage("success", "ค้นหาข้อมูลสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checkProductName)
    if (productid == "") {
      alertMessage("info", "กรอกชื่อสินค้า");
    }
  };

  function HomeIcon(props: any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }

  function BackIcon() {
    return (
      <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 18 18">
        <path d="M15 8.25H5.87l4.19-4.19L9 3 3 9l6 6 1.06-1.06-4.19-4.19H15v-1.5z" />
      </svg>
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
            <Grid item >
              <IconButton
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/SplitsystemManager"
              >
                <BackIcon />
              </IconButton>
            </Grid>
            <Grid item>
              <IconButton
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
              >
                <HomeIcon color="inherit" />
              </IconButton>
            </Grid>

            <Grid item>
              <Button className={classes.button} variant="outlined" color="inherit"
                size="small" component={RouterLink}
                to="/orderproduct">
                ADD DATA
              </Button>
            </Grid>
            <Grid item>

            </Grid>
            <Grid item>
              <IconButton color="inherit" className={classes.iconButtonAvatar}>
                <Avatar src='o' alt={cookieEmail} />
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
                ระบบค้นหารายการสั่งซื้อสินค้าเข้ามาในคลัง
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
          <Tab textColor="inherit" label="Search Data" />
        </Tabs>

      </AppBar>


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
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อสินค้าที่ต้องการค้นหา
              </Typography>
            </Grid>
            <Grid item xs={4}>
              <FormControl
                fullWidth
                variant="outlined"

              ><TextField
                  id="product"
                  label="product"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={productid}
                  onChange={Product_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                />
              </FormControl>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}> </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Button

                variant="contained"
                color="primary"
                size="large"
                className={classes.button}
                onClick={() => {
                  getCheckOrderproduct();
                  setSearch(true);
                }}

                startIcon={<SearchIcon
                />}
              >
                Search
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
      <Grid item xs={12}>
        <Paper>
          {search ? (
            <div>
              {  checkProductName ? (
                <TableContainer component={Paper}>
                  <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                      <TableRow>
                        <TableCell align="center">No.</TableCell>
                        <TableCell align="center">Manager</TableCell>
                        <TableCell align="center">Product</TableCell>
                        <TableCell align="center">Typeproduct</TableCell>
                        <TableCell align="center">Company</TableCell>
                        <TableCell align="center">Stock</TableCell>
                        <TableCell align="center">Date</TableCell>
                        <TableCell align="center">Shipment</TableCell>
                        <TableCell align="center">Detail</TableCell>
                        <TableCell align="center">Manage</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {orderproducts.filter((filter: any) =>
                        filter.edges?.product?.nameProduct.startsWith(productid)).map((item: any) => (
                          <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
                            <TableCell align="center">{item.edges?.managers.name}</TableCell>
                            <TableCell align="center">{item.edges?.product.nameProduct}</TableCell>
                            <TableCell align="center">{item.edges?.typeproduct?.typeproduct}</TableCell>
                            <TableCell align="center">{item.edges?.company?.name}</TableCell>
                            <TableCell align="center">{item.stock}</TableCell>
                            <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH:mm:ss')}</TableCell>
                            <TableCell align="center">{item.shipment}</TableCell>
                            <TableCell align="center">{item.detail}</TableCell>

                            <TableCell align="center">
                              <Button
                                onClick={() => {
                                  deleteSystemequipments(item.id);
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
              )
                : productid == "" ? (
                  <div>
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">No.</TableCell>
                            <TableCell align="center">Manager</TableCell>
                            <TableCell align="center">Product</TableCell>
                            <TableCell align="center">Typeproduct</TableCell>
                            <TableCell align="center">Company</TableCell>
                            <TableCell align="center">Stock</TableCell>
                            <TableCell align="center">Date</TableCell>
                            <TableCell align="center">Shipment</TableCell>
                            <TableCell align="center">Detail</TableCell>
                            <TableCell align="center">Manage</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>
                          {orderproducts.filter((filter: any) =>
                            filter.edges?.product?.nameProduct.startsWith(productid)).map((item: any) => (
                              <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.edges?.managers.name}</TableCell>
                                <TableCell align="center">{item.edges?.product.nameProduct}</TableCell>
                                <TableCell align="center">{item.edges?.typeproduct?.typeproduct}</TableCell>
                                <TableCell align="center">{item.edges?.company?.name}</TableCell>
                                <TableCell align="center">{item.stock}</TableCell>
                                <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH:mm:ss')}</TableCell>
                                <TableCell align="center">{item.shipment}</TableCell>
                                <TableCell align="center">{item.detail}</TableCell>

                                <TableCell align="center">
                                  <Button
                                    onClick={() => {
                                      deleteSystemequipments(item.id);
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
                  </div>
                ) : null}
            </div>
          ) : null}
        </Paper>
      </Grid>
    </div>
  );
}
