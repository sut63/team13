import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'
import Orderonline from './components/Orderonline'
import Orderonlinetable from './components/Orderonline/Tableorderonline'
import Orderproduct from './components/orderproduct'
import Stock from './components/Stock'
import Salary from './components/Salary'
import SalaryTable from './components/SalaryTable'
import Tablestock from './components/Tablestock'
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/Orderonline', Orderonline);
    router.registerRoute('/Orderonlinetable', Orderonlinetable);
    router.registerRoute('/Orderproduct', Orderproduct);
    router.registerRoute('/Stock', Stock);
    router.registerRoute('/Salary', Salary);
    router.registerRoute('/SalaryTable', SalaryTable);
    router.registerRoute('/Tablestock', Tablestock);

  },
});
