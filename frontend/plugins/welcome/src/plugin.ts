import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'
import SignInOrderproduct from './components/orderproduct/SignInOrderproduct'
import Orderonline from './components/Orderonline'
import Orderonlinetable from './components/Orderonline/Tableorderonline'
import SignInOrderonline from './components/Orderonline/SignInOrderonline'
import Orderproduct from './components/orderproduct'
import Stock from './components/Stock'
import SplitsystemManager from './components/SplitsystemManager'
import Salary from './components/Salary'
import SalaryTable from './components/SalaryTable'
import Tablestock from './components/Tablestock'
import Tableorderproduct from './components/orderproduct/Tableorderproduct'
import LoginEmployee from './components/Stock/LoginEmployee'
<<<<<<< HEAD
import ManagerPage from './components/ManagerPage'
=======
import Promotion from './components/Promotion';
>>>>>>> 8f45709b2cda52649f23a4b50b57103e054c12ae

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/SignInOrderproduct', SignInOrderproduct);
    router.registerRoute('/Orderonline', Orderonline);
    router.registerRoute('/Orderonlinetable', Orderonlinetable);
    router.registerRoute('/SignInOrderonline', SignInOrderonline);
    router.registerRoute('/Orderproduct', Orderproduct);
    router.registerRoute('/Stock', Stock);
    router.registerRoute('/SplitsystemManager', SplitsystemManager);
    router.registerRoute('/Salary', Salary);
    router.registerRoute('/SalaryTable', SalaryTable);
    router.registerRoute('/Tablestock', Tablestock);
    router.registerRoute('/Tableorderproduct', Tableorderproduct);
    router.registerRoute('/Promotion', Promotion);
    router.registerRoute('/LoginEmployee', LoginEmployee);
    router.registerRoute('/ManagerPage', ManagerPage);
  },
});
