
export default {
  chain: {
    ethereum: {
      url: 'https://rinkeby.infura.io/v3/81ce7d801fd34ee593ff1ac2298607c0',
    },
  },
  log: {
    level: 'debug',
    path: '',
  },
  metrics: {
    host: '',
    port: 9200,
  },
  sdk: {
    depot: 'http://localhost:5002',
    nym: 'http://localhost:5001',
  },
  user: {
    domain: 'domenetwork.io',
    email: 'dustin.engle@domenetwork.io',
    username: 'dustin',
  },
  wallet: {
    path: 'dome-wallet'
  },
}
