import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import StarIcon from '@material-ui/icons/StarBorder';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import { Link as RouterLink } from 'react-router-dom';
import CardMedia from '@material-ui/core/CardMedia';
import CardActionArea from '@material-ui/core/CardActionArea';
import CameraIcon from '@material-ui/icons/PhotoCamera';
import { Cookies } from '../orderproduct/SignInOrderproduct/Cookie'


function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  '@global': {
    ul: {
      margin: 0,
      padding: 0,
      listStyle: 'none',
    },
  },
  appBar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
  },
  toolbar: {
    flexWrap: 'wrap',
  },
  toolbarTitle: {
    flexGrow: 1,
  },
  link: {
    margin: theme.spacing(1, 1.5),
  },
  heroContent: {
    padding: theme.spacing(8, 0, 6),
  },
  cardHeader: {
    backgroundColor:
      theme.palette.type === 'light' ? theme.palette.grey[200] : theme.palette.grey[700],
  },

  card: {
    height: '200%',
    display: 'flex',
    flexDirection: 'column',
  },
  cardMedia: {
    paddingTop: '100%', // 16:9
  },

  cardPricing: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'baseline',
    marginBottom: theme.spacing(2),
  },
  footer: {
    borderTop: `1px solid ${theme.palette.divider}`,
    marginTop: theme.spacing(8),
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3),
    [theme.breakpoints.up('sm')]: {
      paddingTop: theme.spacing(6),
      paddingBottom: theme.spacing(6),
    },
  },
}));

const tiers = [
  {
    image: "https://cutt.ly/IjWi4y3",
    title: 'ผู้จัดการ',
    description: [/*'ระบบย่อย', 'บันทึกเงินเดือนพนักงาน', 'ตารางเวลาทำงานพนักงาน'*/],
    buttonText: 'Contineus',
    buttinLink: "/signinorderproduct",
    buttonVariant: 'contained',
  },
  {
    image: "https://cutt.ly/MjWi5uk",
    title: 'พนักงาน',
    description: [
     /* '50 users included',
      '30 GB of storage',
      'Help center access',
      'Phone & email support',*/
    ],
    buttonText: 'Contineus',
    buttinLink: "/LoginEmployee",
    buttonVariant: 'contained',
  },
  {
    image: "https://cutt.ly/WjWoq99",
    title: 'ลูกค้า',
    description: [
      /*'20 users included',
      '10 GB of storage',
      'Help center access',
      'Priority email support',*/
    ],
    buttonText: 'Contineus',
    buttinLink: "/SignInOrderonline",
    buttonVariant: 'contained',
  },
  
  
];
const footers = [
  {
    title: 'Poommin Phinphimai B6111090',
    description: 'SOURCE CODE',
    description2: 'Orderproduct System',
    description3: 'Contact us',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/orderproduct",
    link2:"https://www.facebook.com/poommin2543/"
  },
  {
    title: 'Poommin Phinphimai',
    description: 'SOURCE CODE',
    description2: 'Orderproduct System',
    description3: 'Contact us',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/orderproduct",
    link2:"https://www.facebook.com/poommin2543/"
  },
  {
    title: 'Poommin Phinphimai',
    description: 'SOURCE CODE',
    description2: 'Orderproduct System',
    description3: 'Contact us',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/orderproduct",
    link2:"https://www.facebook.com/poommin2543/"
  },
  {
    title: 'Poommin Phinphimai',
    description: 'SOURCE CODE',
    description2: 'Orderproduct System',
    description3: 'Contact us',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/orderproduct",
    link2:"https://www.facebook.com/poommin2543/"
  },
  {
    title: 'Poommin Phinphimai',
    description: 'SOURCE CODE',
    description2: 'Orderproduct System',
    description3: 'Contact us',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/orderproduct",
    link2:"https://www.facebook.com/poommin2543/"
  },
  {
    title: 'Poommin Phinphimai',
    description: 'SOURCE CODE',
    description2: 'Orderproduct System',
    description3: 'Contact us',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/orderproduct",
    link2:"https://www.facebook.com/poommin2543/"
  },
  
];


export default function Pricing() {
  const classes = useStyles();
 

  return (
    <React.Fragment>
      <CssBaseline />
      <AppBar position="static" color="default" elevation={0} className={classes.appBar}>
        <Toolbar className={classes.toolbar}>
          <Typography variant="h3" color="inherit" noWrap className={classes.toolbarTitle}>
            ระบบฟาร์มมาร์ท
          </Typography>
          <nav>
            <Link variant="button" color="textPrimary" href="https://github.com/sut63/team13" className={classes.link}>
            source code
            </Link>
            
          </nav>
          <Button href="#" color="primary" variant="outlined" className={classes.link}
          component={RouterLink}
          to="/signinorderproduct">
            
          </Button>
        </Toolbar>
      </AppBar>
      {/* Hero unit */}
      <Container maxWidth="sm" component="main" className={classes.heroContent}>
        <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
          ยินดีต้อนรับ
        </Typography>
        {/*<Typography variant="h5" align="center" color="textSecondary" component="p">
          Quickly build an effective pricing table for your potential customers with this layout.
          It&apos;s built with default Material-UI components with little customization.
  </Typography>*/}
      </Container>
      {/* End hero unit */}
      <Container maxWidth="md" component="main">
        <Grid container spacing={5} alignItems="flex-end">
            {tiers.map((tier) =>  (
              <Grid item key={tier} xs={12} sm={6} md={4}>
                <Card className={classes.card}>
                  <CardMedia
                    className={classes.cardMedia}
                    image={tier.image}
                    title="Image title"
                  />
                  <CardContent className={classes.card}>
                  <ul>
                    {tier.description.map((line) => (
                      <Typography component="li" variant="subtitle1" align="center" key={line}>
                          {tier.title}
                      </Typography>
                    ))}
                  </ul>
                    <Typography>
                    {tier.title}
                    </Typography>
                  </CardContent>
                  <CardActions>
                    <Button size="small" color="primary" to={tier.buttinLink} component={RouterLink}>
                      {tier.buttonText}
                    </Button>

                  </CardActions>
                </Card>

              </Grid>

            ))}
        </Grid>
      </Container>
      
      
      <Container maxWidth="md" component="footer" className={classes.footer}>
        <Grid container spacing={4} justify="space-evenly">
          {footers.map((footer) => (
            <Grid item xs={6} sm={3} key={footer.title}>
              <Typography variant="h6" color="textPrimary" gutterBottom>
                {footer.title}
              </Typography>
              <ul>
                    
                  <li>
                    <Link href= '' variant="subtitle1" color="textSecondary">
                    {footer.description2}
                    </Link>
                  </li>
                  <li>
                    <Link href= {footer.link}variant="subtitle1" color="textSecondary">
                    {footer.description}
                    </Link>
                  </li>
                  <li>
                    <Link href= {footer.link2}variant="subtitle1" color="textSecondary">
                    {footer.description3}
                    </Link>
                  </li>
               
              </ul>
            </Grid>
          ))}
        </Grid>
        <Box mt={5}>
          <Copyright />
        </Box>
      </Container>
      
    </React.Fragment>
  );
}