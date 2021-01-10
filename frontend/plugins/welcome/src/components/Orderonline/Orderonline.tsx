import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { InputLabel, MenuItem, Select } from '@material-ui/core';
import Typography from '@material-ui/core/Typography';
import TableCell from '@material-ui/core/TableCell';
import Avatar from '@material-ui/core/Avatar';
import Box from '@material-ui/core/Box';

import { EntProduct } from '../../api/models/EntProduct';
import { EntTypeproduct } from '../../api/models/EntTypeproduct';
import { EntPaymentchannel } from '../../api/models/EntPaymentchannel';
import { EntCustomer } from '../../api/models/EntCustomer';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
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

export default function Orderonline() {
  const classes = useStyles();
  const profile = { givenName: 'to PositionAssingment' };
  const api = new DefaultApi();

  //const [Positionassingment, setPositionassingment] = React.useState(initialUserState);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [products, setProducts] = React.useState<EntProduct[]>([]);
  const [typeproducts, setTypeproducts] = React.useState<EntTypeproduct[]>([]);
  const [paymentchannels, setPaymentchannels] = React.useState<EntPaymentchannel[]>([]);
  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);

  const [product, setProduct] = useState(Number);
  const [typeduct, setTypeproduct] = useState(Number);
  const [paymentchannel, setPaymentchannel] = useState(Number);
  const [customer, setCustomer] = useState(Number);
  const [stock, setScotk] = useState(Number);
  const [addedtime, Setaddedtime] = useState(String);


  useEffect(() => {
    const getProducts = async () => {
      const p = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(p);
    };
    getProducts();

    const getTypeproducts = async () => {
      const d = await api.listTypeproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setTypeproducts(d);
    };
    getTypeproducts();

    const getPaymentchannels = async () => {
      const pay = await api.listPaymentchannel({ limit: 10, offset: 0 });
      setLoading(false);
      setPaymentchannels(pay);
    };
    getPaymentchannels();

    const getCustomers = async () => {
        const c = await api.listCustomer({ limit: 10, offset: 0 });
        setLoading(false);
        setCustomers(c);
      };
      getCustomers();

  }, [loading]);


  const handletimeChange = (event: any) => {
    Setaddedtime(event.target.value as string);
  };

  const CreateOrderonline = async () => {
    const orderonline = {
      productid: product,
      typeductid: typeduct,
      paymentchannelid: paymentchannel,
      customerid: customer,
      stock: stock,
      addedtime: addedtime + "00:00+07:00"
    }

    const res: any = await api.createOrderonline({ orderonline: orderonline });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };


  const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProduct(event.target.value as number);
  };

  const typeproduct_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTypeproduct(event.target.value as number);
  };

  const paymentchannel_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPaymentchannel(event.target.value as number);
  };

  const customer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCustomer(event.target.value as number);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to PositionAssingment'}`}
        subtitle="Select the Position you want to be in."
      >

        <Avatar>D</Avatar>
        <Typography component="div" variant="body1">
          <Box color="Dang@gmail.com">Dang@gmail.com</Box>
          <Box color="secondary.main"></Box>
        </Typography>

      </Header>
      <Content>

        <ContentHeader title="PositionAssingment">          
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
        </ContentHeader>

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
                  value={daystart}
                  onChange={handletimeChange}
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
                <InputLabel id="physician_id-label">Physician</InputLabel>
                <Select
                  labelId="physician_id-label"
                  label="Physician"
                  id="physician_id"
                  value={userid}
                  onChange={physician_id_handleChange}
                  style={{ width: 300 }}
                >
                  {physicians.map((item: EntPhysician) =>
                    <MenuItem value={item.id}>{item.eMAIL}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="department_id-label">Department</InputLabel>
                <Select
                  labelId="department_id-label"
                  label="Department"
                  id="department_id"
                  value={department}
                  onChange={department_id_handleChange}
                  style={{ width: 300 }}
                >
                  {departments.map((item: EntDepartment) =>
                    <MenuItem value={item.id}>{item.departmentname}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="position_id-label">Position</InputLabel>
                <Select
                  labelId="position_id-label"
                  label="Position"
                  id="position_id"
                  value={position}
                  onChange={position_id_handleChange}
                  style={{ width: 300 }}
                >
                  {positions.map((item: EntPosition) =>
                    <MenuItem value={item.id}>{item.nameposition}</MenuItem>)}
                </Select>
              </FormControl>
            </TableCell>
            <div className={classes.margin}>
              <TableCell align="right">
                <Button
                  onClick={() => {
                    createOrderonline();
                  }}
                  variant="contained"
                  color="primary"
                  style={{ marginLeft: 545, width: 200 }}
                >
                  SAVE DATA
             </Button>
              </TableCell>
              <TableCell align="right">
                <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/main"
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
}
