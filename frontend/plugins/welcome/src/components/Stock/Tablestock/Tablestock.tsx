import React, { useState, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../../api/apis';
import { EntStock } from '../../../api/models/EntStock';
import { FormControl, Grid, IconButton, InputLabel, Menu, SvgIcon, TextField} from '@material-ui/core';
import {
  Button
} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2';
import { Cookies } from '../LoginEmployee/Cookie';
import { EntProduct } from '../../../api';
import { Content, Header, Page, pageTheme } from '@backstage/core';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    table: {
      minWidth: 650,
    },
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(2),
    },
    withoutLabel: {
      marginTop: theme.spacing(2),
    },
    textField: {
      width: '25ch',
    },
  }),
);

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

export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  var ck = new Cookies()
  var cookieEmail = ck.GetCookie()
  var cookieID = ck.GetID()
  var cookieName = ck.GetName()
  const [products, setProducts] = useState<EntProduct[]>([]);
  const [loading, setLoading] = useState(true);
  const [productid, setProductid] = useState(String);

  const [checkProductName, setProductNames] = useState(false);
  const profile = { givenName: '' };
  const [search, setSearch] = useState(false);
  //const [checkvalue, setCheckValue] = useState(false);
  const [stocks, setStocks] = useState<EntStock[]>([]);

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getProducts = async () => {
      const pr = await api.listProduct({ offset: 0 });
      setLoading(false);
      setProducts(pr);
    };
    getProducts();
   
    const getStockproduct = async () => {
      const sp = await api.listStock({ offset: 0 });
      setLoading(false);
      setStocks(sp);
    };
    getStockproduct();
  }, [loading]);

  const Product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProductid(event.target.value as string);
    setProductNames(false);
    setSearch(false);
  };

  const getCheckstock = async () => {
    var check = false;
    stocks.map(item => {
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
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  function HomeIcon(props) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ระบบค้นหา  ${profile.givenName || 'Stock สินค้าหน้าร้าน'}`}
      //subtitle="ทำการค้นหา Stock โดยใส่ชื่อสินค้าลงไป"
      >
        <div style={{ marginLeft: 300 }}><h3>{cookieName}</h3></div>
        <IconButton
          aria-label="account of current patientofphysician"
          aria-controls="menu-appbar"
          aria-haspopup="true"
          onClick={handleMenu}
        >
          <AccountCircle />
        </IconButton>
        <Menu
          id="menu-appbar"
          anchorEl={anchorEl}
          anchorOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          keepMounted
          transformOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          open={open}
          onClose={handleClose}
        >
          <Button className={classes.button} variant="outlined" color="inherit"
            size="small" component={RouterLink}
            to="/LoginEmployee">
            logout
              </Button>
        </Menu>
        <Button
          variant="outlined"
          color="inherit"
          size="small"
          className={classes.button}
          startIcon={<HomeIcon />}
          component={RouterLink}
          to="/Stock"
        >
          AddStock
         </Button>
      </Header>
      <Content>
        <div className={classes.root}>
          <form noValidate autoComplete="off"></form>
          <Grid container alignItems="center" spacing={4}>
            <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>   <center>
              <h3 align='right'>ชื่อพนักงาน</h3></center>
            </Grid>
            <Grid item xs={4}>
              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
              ><InputLabel></InputLabel>
                <div style={{ marginRight: 300 }}><h4>{cookieName}</h4></div>
              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>
            <Grid item xs={4}><center>
              <h3 align='right'>ชื่อสินค้า</h3></center>
            </Grid>
            <Grid item xs={4}>
              <FormControl
                fullWidth
                className={classes.margin}
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
            <Grid item xs={4}>
            </Grid>
            <Grid item xs={4}>
            </Grid>
            <Grid item xs={4}>
              <Button
                onClick={() => {
                  getCheckstock();
                  setSearch(true);
                }}
                variant="outlined"
                color="inherit"
                size="large"
                className={classes.button}
                startIcon={<SearchIcon />}
              >
                Search
              </Button>
            </Grid>
            <Grid item xs={4}></Grid>

            <Grid container justify="center"></Grid>
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
                              <TableCell align="center">Employee</TableCell>
                              <TableCell align="center">IDStock</TableCell>
                              <TableCell align="center">Typeproduct</TableCell>
                              <TableCell align="center">Product</TableCell>
                              <TableCell align="center">Zoneproduct</TableCell>
                              <TableCell align="center">Priceproduct</TableCell>
                              <TableCell align="center">Amount</TableCell>
                              <TableCell align="center">Timesave</TableCell>
                            </TableRow>
                          </TableHead>
                          <TableBody>
                            {stocks.filter((filter: any) => filter.edges?.product?.nameProduct.startsWith(productid)).map((item: any) => (
                              <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.edges?.employee?.name}</TableCell>
                                <TableCell align="center">{item.iDstock}</TableCell>
                                <TableCell align="center">{item.edges?.typeproduct?.typeproduct}</TableCell>
                                <TableCell align="center">{item.edges?.product?.nameProduct}</TableCell>
                                <TableCell align="center">{item.edges?.zoneproduct?.zone}</TableCell>
                                <TableCell align="center">{item.priceproduct}</TableCell>
                                <TableCell align="center">{item.amount}</TableCell>
                                <TableCell align="center">{item.time}</TableCell>
                                <TableCell align="center">
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
                                  <TableCell align="center">Employee</TableCell>
                                  <TableCell align="center">IDStock</TableCell>
                                  <TableCell align="center">Typeproduct</TableCell>
                                  <TableCell align="center">Product</TableCell>
                                  <TableCell align="center">Zoneproduct</TableCell>
                                  <TableCell align="center">Priceproduct</TableCell>
                                  <TableCell align="center">Amount</TableCell>
                                  <TableCell align="center">Timesave</TableCell>
                                </TableRow>
                              </TableHead>
                              <TableBody>
                                {stocks.filter((filter: any) => filter.edges?.product?.nameProduct.startsWith(productid)).map((item: any) => (
                                  <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.employee?.name}</TableCell>
                                    <TableCell align="center">{item.iDstock}</TableCell>
                                    <TableCell align="center">{item.edges?.typeproduct?.typeproduct}</TableCell>
                                    <TableCell align="center">{item.edges?.product?.nameProduct}</TableCell>
                                    <TableCell align="center">{item.edges?.zoneproduct?.zone}</TableCell>
                                    <TableCell align="center">{item.priceproduct}</TableCell>
                                    <TableCell align="center">{item.amount}</TableCell>
                                    <TableCell align="center">{item.time}</TableCell>
                                    <TableCell align="center"></TableCell>
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
          </Grid>
        </div>
      </Content>
    </Page>
  );
}
