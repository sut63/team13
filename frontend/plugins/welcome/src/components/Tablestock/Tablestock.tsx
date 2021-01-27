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
import { DefaultApi } from '../../api/apis';
import { EntStock } from '../../api/models/EntStock';
import { AppBar, IconButton, Menu, Toolbar, Typography } from '@material-ui/core';
import {
  MenuItem, 
  Button, 
} from '@material-ui/core';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2';
import { Cookies } from '../Stock/LoginEmployee/Cookie';
import { EntProduct } from '../../api';

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

 const [stocks, setStocks] = useState<EntStock[]>();
 
 let employeeID = Number(cookieID)
 let productID = Number(productid)
 console.log(employeeID)
 
 useEffect(() => {
   const getStocks = async () => {
     const res = await api.listStock({ limit: 10, offset: 0 });
     setLoading(false);
     setStocks(res);
   };
   getStocks();
 }, [loading]);
 
 const deleteStocks = async (id: number) => {
   const res = await api.deleteStock({ id: id });
   setLoading(true);
 };
 const orderproduct = {
  managerID,
  productID,
}
console.log(orderproduct)

const Product_id_handleChange = (event: any) => {
  setProductid(event.target.value);
}
var lenstock: number

const getCheckinsorder = async () => {
  const res = await api.getStock({ id: productid })
  setProducts(res)
  lenstock = res.length
  if (lenstock > 0) {
    //setOpen(true)
    Toast.fire({
      icon: 'success',
      title: 'ค้นหาข้อมูลสำเร็จ',
    })
  } else {
    //setFail(true)
    Toast.fire({
      icon: 'error',
      title: 'ค้นหาข้อมูลไม่สำเร็จ',
    })
  }
}
 
 return (
 
   
  
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
           <Button
              style={{ marginLeft: 20 ,width : 100 }}
              component={RouterLink}
              to="/Stock"
              variant="contained"
            >
              Back
             </Button>
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

             <Button
              onClick={() => {
                deleteStocks(item.id);
              }}
        variant="contained"
        color="secondary"
        className={classes.button}
        startIcon={<DeleteIcon />}
      >
        Delete
      </Button>





             </TableCell>
             
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
