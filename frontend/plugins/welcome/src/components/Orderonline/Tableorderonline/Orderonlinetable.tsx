import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Link as RouterLink } from 'react-router-dom';
import {
  Header,
  Page,
  pageTheme,
} from '@backstage/core';

import { DefaultApi } from '../../../api/apis';
import { EntOrderonline } from '../../../api/models/EntOrderonline';
import {
  Avatar,
  Button, IconButton,
} from '@material-ui/core';
import { Cookiesonline } from '../SignInOrderonline/Cookie'
import moment from 'moment';


const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
  iconButtonAvatar: {
    padding: 4,
  },
});

export default function Orderonlinetable() {
  var ck = new Cookiesonline()
  var cookieEmail = ck.GetCookie()
  var cookieID = ck.GetID()
  var cookieName = ck.GetName()
  const classes = useStyles();
  const api = new DefaultApi();

  const [orderonlines, setOrderonlines] = useState<EntOrderonline[]>();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getOrderonlines = async () => {
      const res = await api.listOrderonline({ limit: 10, offset: 0 });
      setLoading(false);
      setOrderonlines(res);
    };
    getOrderonlines();
  }, [loading]);

  const deleteOrderonlines = async (id: number) => {
    const res = await api.deleteOrderonline({ id: id });
    setLoading(true);
  };

  return (

    <Page theme={pageTheme.home}>
      <Header
        title=" Welcome to Orderonline "
        subtitle="Select Product you want to be in."
      >

        <TableCell align="right">
          <Button
            style={{ marginLeft: 1 }}
            component={RouterLink}
            to="/SearchOrderonline"
            variant="contained"
          >
            Back
             </Button>
        </TableCell>

        <IconButton color="inherit" className={classes.iconButtonAvatar}>
                <Avatar src='o' alt={cookieEmail} />
        </IconButton>

      </Header>


      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">No.</TableCell>
              <TableCell align="center">Customer</TableCell>
              <TableCell align="center">Typeproduct</TableCell>
              <TableCell align="center">Product</TableCell>
              <TableCell align="center">Stock</TableCell>
              <TableCell align="center">Paymentchannel</TableCell>
              <TableCell align="center">Accountnumber</TableCell>
              <TableCell align="center">CVV</TableCell>
              <TableCell align="center">Time</TableCell>
              <TableCell align="center">Amount</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {orderonlines === undefined
              ? null
              : orderonlines.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.customer?.name}</TableCell>
                  <TableCell align="center">{item.edges?.typeproduct?.typeproduct}</TableCell>
                  <TableCell align="center">{item.edges?.product?.nameProduct}</TableCell>
                  <TableCell align="center">{item.stock}</TableCell>
                  <TableCell align="center">{item.edges?.paymentchannel?.bank}</TableCell>
                  <TableCell align="center">{item.accountnumber}</TableCell>
                  <TableCell align="center">{item.cvv}</TableCell>
                  <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH:mm')}</TableCell>
                  <TableCell align="center">{item.amount}</TableCell>

                  <TableCell align="center">
                    <Button
                      onClick={() => {
                        deleteOrderonlines(item.id);
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
    </Page>
  );
}
