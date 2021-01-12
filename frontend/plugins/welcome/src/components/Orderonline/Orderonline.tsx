import React, { FC, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';

import { DefaultApi } from '../../api/apis';
import { InputLabel, MenuItem, Select } from '@material-ui/core';
import Typography from '@material-ui/core/Typography';
import TableCell from '@material-ui/core/TableCell';
import Avatar from '@material-ui/core/Avatar';
import Box from '@material-ui/core/Box';
import Swal from 'sweetalert2'; // alert

import { EntProduct } from '../../api/models/EntProduct';
import { EntTypeproduct } from '../../api/models/EntTypeproduct';
import { EntPaymentchannel } from '../../api/models/EntPaymentchannel';
import { EntCustomer } from '../../api/models/EntCustomer';

import SaveIcon from '@material-ui/icons/Save'; // icon save

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },
    withoutLabel: {
      marginTop: theme.spacing(2),
    },
    textField: {
      width: '25ch',
    },
  }),
);

interface orderonline {
  productid: number;
  typeductid: number;
  paymentchannelid: number;
  customerid: number;
  stock: number;
  addedtime: Date;
  // create_by: number;
}

const OrderOnline: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'to Order Online' };
  const http = new DefaultApi();

  const [orderonlines, setOrderonlines] = React.useState<Partial<orderonline>>({});

  const [products, setProducts] = React.useState<EntProduct[]>([]);
  const [typeproducts, setTypeproducts] = React.useState<EntTypeproduct[]>([]);
  const [paymentchannels, setPaymentchannels] = React.useState<EntPaymentchannel[]>([]);
  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);

  const getProducts = async () => {
    const p = await http.listProduct({ limit: 10, offset: 0 });
    setProducts(p);
  };


  const getTypeproducts = async () => {
    const d = await http.listTypeproduct({ limit: 10, offset: 0 });
    setTypeproducts(d);
  };


  const getPaymentchannels = async () => {
    const pay = await http.listPaymentchannel({ limit: 10, offset: 0 });
    setPaymentchannels(pay);
  };


  const getCustomers = async () => {
    const c = await http.listCustomer({ limit: 10, offset: 0 });
    setCustomers(c);
  };

// Lifecycle Hooks
useEffect(() => {
  getProducts();
  getTypeproducts();
  getPaymentchannels();
  getCustomers();
}, []);


  // set data to object orderonline
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof OrderOnline;
    const { value } = event.target;
    setOrderonlines({ ...orderonlines, [name]: value });
    console.log(orderonlines);
  };


  // alert setting
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

  function clear() {
    setOrderonlines({});
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/orderonlines';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(orderonlines),
    };

    console.log(orderonlines); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Orderonline'}`}
        subtitle="Select Product you want to be in."
      >

        <Avatar>D</Avatar>
        <Typography component="div" variant="body1">
          <Box color="Dang@gmail.com">Dang@gmail.com</Box>
          <Box color="secondary.main"></Box>
        </Typography>

      </Header>
      <Content>

        <div className={classes.root}>
          <form noValidate autoComplete="off">



            <TableCell align="left">

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <TextField
                  id="DayStart"
                  label="DayStart"
                  type="date"
                  value={orderonlines.addedtime}
                  onChange={handleChange}
                  //defaultValue="2020-05-24"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>


              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}

              >
                <InputLabel id="product_id-label">Product</InputLabel>
                <Select
                  labelId="product_id-label"
                  label="Product"
                  id="product_id"
                  value={orderonlines.productid}
                  onChange={handleChange}
                  style={{ width: 300 }}
                >
                  {products.map((item: EntProduct) =>
                    <MenuItem value={item.id}>{item.nameProduct}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="typeproduct_id-label">Typeproduct</InputLabel>
                <Select
                  labelId="typeproduct_id-label"
                  label="Typeproduct"
                  id="typeproduct_id"
                  value={orderonlines.typeductid}
                  onChange={handleChange}
                  style={{ width: 300 }}
                >
                  {typeproducts.map((item: EntTypeproduct) =>
                    <MenuItem value={item.id}>{item.typeproduct}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="paymentchannel_id-label">Paymentchannel</InputLabel>
                <Select
                  labelId="paymentchannel_id-label"
                  label="Paymentchannel"
                  id="paymentchannel_id"
                  value={orderonlines.paymentchannelid}
                  onChange={handleChange}
                  style={{ width: 300 }}
                >
                  {paymentchannels.map((item: EntPaymentchannel) =>
                    <MenuItem value={item.id}>{item.bank}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="customer_id-label">Customer</InputLabel>
                <Select
                  labelId="customer_id-label"
                  label="Customer"
                  id="customer_id"
                  value={orderonlines.customerid}
                  onChange={handleChange}
                  style={{ width: 300 }}
                >
                  {customers.map((item: EntCustomer) =>
                    <MenuItem value={item.id}>{item.name}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 302 }}
              >
                <TextField value={orderonlines.stock} id="outlined-number" type='number' InputLabelProps={{
                  shrink: true,
                }} label="กรุณาใส่จำนวน" variant="outlined"
                  onChange={handleChange}
                />
              </FormControl>


            </TableCell>

            <div className={classes.margin}>
              <TableCell align="right">
                <Button
                  variant="contained"
                  color="primary"
                  size="large"
                  startIcon={<SaveIcon />}
                  onClick={save}
                  style={{ marginLeft: 545, width: 200 }}
                >
                  บันทึกการดู
              </Button>
              </TableCell>

              <TableCell align="right">
                <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/"
                  variant="contained"
                >
                  Back
             </Button>
              </TableCell>

            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
};

export default OrderOnline;
