import React, { FC, useState, useEffect } from 'react';
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
import { EntZoneproduct } from '../../api/models/EntZoneproduct';
import { EntEmployee } from '../../api/models/EntEmployee';


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

const Stock: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'Stock Product' };
  
    const api = new DefaultApi();
    const [products, setProducts] = useState<EntProduct[]>([]);
    const [typeproducts, setTypeproducts] = useState<EntTypeproduct[]>([]);
    const [employees, setEmployees] = useState<EntEmployee[]>([]);
    const [zoneproducts, setZoneproducts] = useState<EntZoneproduct[]>([]);
    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [loading, setLoading] = useState(true);
  
    const [productid, setProductid] = useState(Number);
    const [priceproduct, setPriceproduct] = useState(Number);
    const [amount, setAmount] = useState(Number);
    const [time, setTime] = useState(String);
    const [typeproductid, setTypeproductid] = useState(Number);
    const [employeeid, setEmployeeid] = useState(Number);
    const [zoneproductid, setZoneproductid] = useState(Number);
  
    useEffect(() => {
      const getProducts = async () => {
        const pr = await api.listProduct({ limit: 10, offset: 0 });
        setLoading(false);
        setProducts(pr);
      };
      getProducts();
  
  
      const getTypeproducts = async () => {
        const ty = await api.listTypeproduct({ limit: 10, offset: 0 });
        setLoading(false);
        setTypeproducts(ty);
      };
      getTypeproducts();
  
      const getEmployees = async () => {
        const em = await api.listEmployee({ limit: 10, offset: 0 });
        setLoading(false);
        setEmployees(em);
      };
      getEmployees();

  
      const getZoneproducts = async () => {
        const zo = await api.listZoneproduct({ limit: 10, offset: 0 });
        setLoading(false);
        setZoneproducts(zo);
      };
      getZoneproducts();

    }, [loading]);
  
    const handletimeChange = (event: any) => {
        setTime(event.target.value as string);
      };
  
    const CreateStock = async () => {
      const stock = {
        product: productid,
        typeproduct: typeproductid,
        employee: employeeid,
        zoneproduct: zoneproductid,
        priceproduct : priceproduct,
        amount : amount,
        time: time + "00:00+07:00"
  
  
  
      }
      console.log(stock);
  
      const res: any = await api.createStock({ stock: stock });
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
      setProductid(event.target.value as number);
    };
    const employee_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setEmployeeid(event.target.value as number);
    };
  
    const typeproduct_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setTypeproductid(event.target.value as number);
    };

    const zoneproduct_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setZoneproductid(event.target.value as number);
      };

    const amount_id_handleChange = (event: any) => {
        setAmount(event.target.value);
       };

       const priceproduct_id_handleChange = (event: any) => {
        setPriceproduct(event.target.value);
      };

  
       return (
        <Page theme={pageTheme.home}>
          <Header
            title={`Welcome ${profile.givenName || 'to Stock'}`}
            subtitle="Product in stock."
          >
    
            <Avatar>D</Avatar>
            <Typography component="div" variant="body1">
              <Box color="Pang@gmail.com">PANG@gmail.com</Box>
              <Box color="secondary.main"></Box>
            </Typography>
    
          </Header>
          <Content>
    
            <ContentHeader title="Stock Product">          
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
                    <InputLabel id="employee_id-label">Employee</InputLabel>
                    <Select
                      labelId="employee_id-label"
                      label="employee"
                      id="employee_id"
                      value={employeeid}
                      onChange={employee_id_handleChange}
                      style={{ width: 300 }}
                    >
                      {employees.map((item: EntEmployee) =>
                        <MenuItem value={item.id}>{item.name}</MenuItem>)}
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
                      value={typeproductid}
                      onChange={typeproduct_id_handleChange}
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
                    <InputLabel id="product_id-label">Product</InputLabel>
                    <Select
                      labelId="product_id-label"
                      label="Product"
                      id="product_id"
                      value={productid}
                      onChange={product_id_handleChange}
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
                    <InputLabel id="zoneproduct_id-label">Zoneproduct</InputLabel>
                    <Select
                      labelId="zoneproduct_id-label"
                      label="zoneproduct"
                      id="zoneproduct_id"
                      value={zoneproductid}
                      onChange={zoneproduct_id_handleChange}
                      style={{ width: 300 }}
                    >
                      {zoneproducts.map((item: EntZoneproduct) =>
                        <MenuItem value={item.id}>{item.zone}</MenuItem>)}
                    </Select>
                  </FormControl>

                  <FormControl
                    fullWidth
                    className={classes.margin}
                    variant="outlined"

                 ><TextField
                       id="priceproduct"
                      label="Price"
                      variant="outlined"
                      type="number"
                      size="medium"
                      value={priceproduct}
                      onChange={priceproduct_id_handleChange}
                      style={{ width: 300 }}
                 />
                  </FormControl>
    
                  
    
                  <FormControl
                    fullWidth
                    className={classes.margin}
                    variant="outlined"
                    style={{ marginLeft: 560, width: 302 }}
                  >
                    <TextField id="outlined-number" type='number'  InputLabelProps={{
                      shrink: true,}}label="Amount" variant="outlined"
                      onChange = {amount_id_handleChange}
                      />
                  </FormControl>

                  <FormControl
                    fullWidth
                    className={classes.margin}
                    variant="outlined"
                    style={{ marginLeft: 560, width: 600 }}
                  >
                    <TextField
                      id="time"
                      label="time"
                      type="date"
                      value={time}
                      onChange={handletimeChange}
                      //defaultValue="2020-05-24"
                      className={classes.textField}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
                  </FormControl>
    
    
                </TableCell>
    
                <div className={classes.margin}>
                  <TableCell align="right">
                    <Button
                      onClick={() => {
                        CreateStock();
                      }}
                      variant="contained"
                      color="primary"
                      style={{ marginLeft: 545, width: 200 }}
                    >
                      SAVE 
                 </Button>
                  </TableCell>
    
                  <TableCell align="right">
                    <Button
                      style={{ marginLeft: 1 }}
                      component={RouterLink}
                      to="/WelcomePage"
                      variant="contained"
                    >
                      BACK
                 </Button>
                  </TableCell>
                  
                </div>
              </form>
            </div>
          </Content>
        </Page>
      );
    };
    
    export default Stock;