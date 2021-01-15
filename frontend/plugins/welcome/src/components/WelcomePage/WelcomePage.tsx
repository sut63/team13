import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import { Link as RouterLink } from 'react-router-dom';
import CardMedia from '@material-ui/core/CardMedia';


function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Farmmart Website
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
    title: 'Manager',
    description: [/*'ระบบย่อย', 'บันทึกเงินเดือนพนักงาน', 'ตารางเวลาทำงานพนักงาน'*/],
    buttonText: 'Contineus',
    buttinLink: "/signinorderproduct",
    buttonVariant: 'contained',
  },
  {
    image: "https://cutt.ly/MjWi5uk",
    title: 'Employee',
    description: [],
    buttonText: 'Contineus',
    buttinLink: "/LoginEmployee",
    buttonVariant: 'contained',
  },
  {
    image: "https://cutt.ly/WjWoq99",
    title: 'Customer',
    description: [],
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
    description3: 'Contact ',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/orderproduct",
    link2: "https://www.facebook.com/poommin2543/",
    link3: "/orderproduct"
  },
  {
    title: 'Konrawit Kongsri B6118310',
    description: 'SOURCE CODE',
    description2: 'Orderonline System',
    description3: 'Contact ',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/Orderonline",
    link2: "https://www.facebook.com/peerat.kongsree",
    link3: "/Orderonline"
  },
  {
    title: 'Prakaifa Kummungkun B6104245',
    description: 'SOURCE CODE',
    description2: 'StockProduct    System',
    description3: 'Contact ',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/Stock",
    link2: "https://www.facebook.com/prakaifa.kummungkum/",
    link3: "/Stock"
  },
  {
    title: 'Nontakorn Payusuk B6111052',
    description: 'SOURCE CODE',
    description2: 'Salary  System',
    description3: 'Contact ',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/Salary",
    link2: "https://www.facebook.com/profile.php?id=100002730288880",
    link3: "/Salary"
  },
  {
    title: 'Thanabodee Petchrey B6118631',
    description: 'SOURCE CODE',
    description2: 'EmployeeWorkingHours',
    description3: 'Contact ',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/EmployeeWorkingHours",
    link2: "https://www.facebook.com/profile.php?id=100004791875262",
    link3: "/EmployeeWorkingHours"
  },
  {
    title: 'Pichai Somasa B6116637',
    description: 'SOURCE CODE',
    description2: 'Promotion System',
    description3: 'Contact ',
    link: "https://github.com/sut63/team13/tree/main/frontend/plugins/welcome/src/components/Promotion",
    link2: "https://www.facebook.com/bank.kub.775/",
    link3: "/Promotion"
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
            Farmmart System
          </Typography>
          <nav>
            <Link variant="button" color="textPrimary" href="https://github.com/sut63/team13" className={classes.link}>
              source code
            </Link>
          </nav>
        </Toolbar>
      </AppBar>
      {/* Hero unit */}
      <Container maxWidth="sm" component="main" className={classes.heroContent}>
        <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
          Welcome
        </Typography>

      </Container>
      {/* End hero unit */}
      <Container maxWidth="md" component="main">
        <Grid container spacing={5} alignItems="flex-end">
          {tiers.map((tier) => (
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
                  <Link href={footer.link3} variant="subtitle1" color="textSecondary">
                    {footer.description2}
                  </Link>
                </li>
                <li>
                  <Link href={footer.link} variant="subtitle1" color="textSecondary">
                    {footer.description}
                  </Link>
                </li>
                <li>
                  <Link href={footer.link2} variant="subtitle1" color="textSecondary">
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