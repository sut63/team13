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
import { EntPromotion } from '../../../api/models/EntPromotion';
import {
  Button,
} from '@material-ui/core';



const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function Promotiontable() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [promotion, setPromotion] = useState<EntPromotion[]>();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getPromotion = async () => {
      const res = await api.listPromotion({ limit: 10, offset: 0 });
      setLoading(false);
      setPromotion(res);
    };
    getPromotion();
  }, [loading]);

  const deletePromotion = async (id: number) => {
    const res = await api.deletePromotion({ id: id });
    setLoading(true);
  };

  console.log(promotion)
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
            to="/Promotion"
            variant="contained"
          >
            Back
             </Button>
        </TableCell>

      </Header>


      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">No.</TableCell>
              <TableCell align="center">ชื่อโปรโมชั่น</TableCell>
              <TableCell align="center">สินค้า</TableCell>
              <TableCell align="center">ของแถม</TableCell>
              <TableCell align="center">ส่วนลด</TableCell>
              <TableCell align="center">ราคา</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {promotion === undefined
              ? null
              : promotion.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.promotionName}</TableCell>
                  <TableCell align="center">{item.edges?.product?.nameProduct}</TableCell>
                  <TableCell align="center">{item.edges?.give?.giveawayName}</TableCell>
                  <TableCell align="center">{item.edges?.sale?.sale}</TableCell>
                  <TableCell align="center">{item.price}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                        deletePromotion(item.id);
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
