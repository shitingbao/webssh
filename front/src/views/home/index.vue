<template>
  <div>
    <input v-model="sn" />
    <button @click="homeCliek">client</button>
  </div>
</template>
<script>
export default {
  name: "App",
  mounted() {},
  methods: {
    homeCliek() {
      this.$router.push({ name: "term", query: { sn: this.sn } });
      console.log("push click", this.sn);
    },
  },
  data() {
    return {
      sn: "input your sn",
    };
  },
};
</script>
<style scoped>
.home {
  color: rgb(255, 255, 255);
  /* background: rgb(211, 24, 24); */
}
</style>
