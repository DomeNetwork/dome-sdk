
<script>

const denoms = [
  [1000000000000000000, 'Ether'],
  [1000000000000000, 'Milliether'],
  [1000000000000, 'Microether'],
  [1000000000, 'Gwei'],
  [1000000, 'Mwei'],
  [1000, 'Kwei'],
  [1, 'Wei'],
]

export default {
  computed: {
    amount() {
      for (const denom of denoms) {
        if (this.v >= denom[0]) {
          this.denom = denom[1]
          return (this.v / denom[0]).toFixed(8)
        }
      }
    },
  },
  data: () => ({
    denom: 'Wei',
  }),
  props: {
    v: Number,
  },
}
</script>

<template>
  {{ amount }} {{ denom }}
</template>

<style>
</style>
