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
      sn: "XB01-X001-B003-C888",
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
