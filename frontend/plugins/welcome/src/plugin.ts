import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'
import Orderonline from './components/Orderonline'
import Orderproduct from './components/orderproduct'
import Salary from './components/Salary'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/Orderonline', Orderonline);
    router.registerRoute('/Orderproduct', Orderproduct);
    router.registerRoute('/Salary', Salary);
  },
});
