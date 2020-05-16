<template>
  <div class="container"></div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Home",
  components: {},
  data() {
    return {
      res: null,
    };
  },
  created() {},
  methods: {
    hoge() {
      let _self = this;
      this.$axios.get("http://localhost:8082/sample").then((response) => {
        _self.res = response.data;
        console.log(_self.res);
      });
    },
  },
};
</script>
