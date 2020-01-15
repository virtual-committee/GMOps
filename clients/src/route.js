import About from '@/components/about';
import Home from '@/components/home';
export default [{
    path: '/home',
    component: Home
  },
  {
    path: '/about',
    component: About
  },
  {
    path: '*',
    redirect: '/home'
  }
]