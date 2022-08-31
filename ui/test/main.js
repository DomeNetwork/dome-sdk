
const config = {
  coin: {
    ethereum: {
      url: 'https://rinkeby.infura.io/v3/81ce7d801fd34ee593ff1ac2298607c0',
      ws: 'wss://rinkeby.infura.io/ws/v3/81ce7d801fd34ee593ff1ac2298607c0',
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
const secret = '3333'
const sendToAddr = '0x'

const blob = {
  data: '00010203040506070809',
  metadata: 'Something here!',
  name: 'Test1',
}
const user = {
  avatar: '',
  domain: config.user.domain,
  email: config.user.email,
  password: 'password',
  publicKey: '',
  username: config.user.username,
}

function log(k, v) {
  console.log(`${k}:`, v)
}

async function setupJS() {
  if (!window.dome) {
    setTimeout(setupJS, 1000)
    return
  }

  // Setup
  console.log('config:', config)
  await dome.config(JSON.stringify(config))

  // Wallet
  const isFound = await dome.check()
  if (isFound) {
    await dome.open(secret)
  } else {
    const words = await dome.mnemonic()
    log('words', words)

    await dome.load(words)

    await dome.save(secret)
  }

  user.publicKey = await dome.publicKey()

  // Nym
  // await dome.register(JSON.stringify(user))

  // await dome.login(JSON.stringify(user))

  // for (let i = 0; i <= 5; i++) {
  //   const acct = await dome.generate('Ethereum')
  //   log('account', acct)
  //   await dome.key(JSON.stringify({
  //     name: acct.address,
  //     publicKey: acct.publicKey,
  //     used: false,
  //   }))
  // }

  // const handles = await dome.search(JSON.stringify({ term: 'd' }))
  // log('handles', handles)

  // const lookup = await dome.lookup(JSON.stringify(user))
  // log('lookup', lookup)

  // await dome.update(JSON.stringify(user))

  // await dome.forgot(JSON.stringify(user))

  // await dome.reset(otp, JSON.stringify(user))

  // Depot
  // await dome.upload(JSON.stringify(blob))

  // await dome.metadata(JSON.stringify(blob))

  // await dome.download(JSON.stringify(blob))

  // Ethereum
  const coins = await dome.coins()
  log('coins', coins)

  const eth = coins.find(coin => coin.name === 'Ethereum')
  log('eth', eth)

  const accts = await dome.accounts(eth.name)
  log('accts', accts)

  const balance = await dome.balance(eth.name)
  log('balance', balance)

  const gas = await dome.gas(eth.name)
  log('gas', gas)

  const listener = (v) => console.log(v)
  await dome.subscribe(eth.name, listener)

  setTimeout(async () => {
    await dome.unsubscribe(eth.name)
    log('unsubscribe')
  }, 30000)

  // // TODO: setup amount and send to address
  // // const tx = await dome.send(eth, sendToAddr, amount, data)
  // // log('tx', tx)
}
