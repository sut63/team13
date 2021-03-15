import React, { useState, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import DeleteIcon from '@material-ui/icons/Delete';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../../api/apis';
import { EntStock } from '../../../api/models/EntStock';
import { AppBar, FormControl, Grid, IconButton, InputLabel, Menu, Select, SvgIcon, TextField, Toolbar, Typography } from '@material-ui/core';
import {
  MenuItem, 
  Button, 
} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import { Alert } from '@material-ui/lab';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2';
import { Cookies } from '../LoginEmployee/Cookie';
import { EntProduct } from '../../../api';
import { Content, ContentHeader, Header, Page, pageTheme } from '@backstage/core';


const useStyles = makeStyles((theme: Theme) =>
createStyles({
  table: {
    minWidth :650,
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



/*const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});*/
 //alert
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
 const [productid, setProductid] = useState(Number);
 const [alert, setAlert] = useState(true);
 const profile = { givenName: '' };

 const [stocks, setStocks] = useState<EntStock[]>();
 
 let employeeID = Number(cookieID)
 let productID = Number(productid)
 console.log(employeeID)
 
 useEffect(() => {
   const getProducts = async () => {
     const pr = await api.listProduct({ limit: 10, offset: 0 });
     setLoading(false);
     setProducts(pr);
   };
   getProducts();
 }, [loading]);
 
 const deleteStocks = async (id: number) => {
   const res = await api.deleteStock({ id: id });
   setLoading(true);
 };

 const stock = {
  employeeID,
  productID,
}
console.log(stock)

}

const getCheckstock = async () => {
  const res = await api.getStock({ id: productid })
  setStocks(res)
  lenstock = res.length
  if (lenstock > 0) {
    //setOpen(true)
    Toast.fire({
      icon: 'success',
      title: 'มีสินค้าใน Stock ค้นหาสำเร็จ',
    })
  } else {
    //setFail(true)
    Toast.fire({
      icon: 'error',
      title: 'ไม่มีสินค้าใน Stock ค้นหาไม่สำเร็จ',
    })
  }
}

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
           startIcon={<HomeIcon  />}
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

                {/*} <Select
                labelId="employee_id-label"
                label="employee"
                id="eployee_id"
                value={employeeid}
                onChange={employee_id_handleChange}
                style={{ marginRight : 300 ,width: 300 }}
              >
                {employees.map((item: EntEmployee) =>
                  <MenuItem value={item.id}>{item.name}</MenuItem>)}
                </Select>*/}
                <div style={{ marginRight: 300 }}><h4>{cookieName}</h4></div>

              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>





            
            <Grid item xs={4}><center>
              <h3 align='right'>สินค้า</h3></center>
            </Grid>
            <Grid item xs={4}>
              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
              ><InputLabel>Product</InputLabel>
                <Select
                  labelId="product_id-label"
                  label="product"
                  id="product_id"
                  value={productid}
                  onChange={Product_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                >
                  {products.map((item: EntProduct) =>
                    <MenuItem key={item.id} value={item.id}>{item.nameProduct}</MenuItem>)}
                </Select>
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


           









   
  
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Employee</TableCell>
           <TableCell align="center">IDEmployee</TableCell>
           <TableCell align="center">Typeproduct</TableCell>
           <TableCell align="center">Product</TableCell>
           <TableCell align="center">Zoneproduct</TableCell>
           <TableCell align="center">Priceproduct</TableCell>
           <TableCell align="center">Amount</TableCell>
           <TableCell align="center">Time</TableCell>
           
         </TableRow>
       </TableHead>
       <TableBody>
       {stocks === undefined 
          ? null
          : stocks.map((item : any)=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.employee?.name}</TableCell>
             <TableCell align="center">{item.iDcardemployee}</TableCell>
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

   <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>
          </Grid>

        </div>


      </Content>
    </Page>

 );
}
