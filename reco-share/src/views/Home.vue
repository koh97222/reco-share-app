<template>
  <div class="home">
    <b-button @click="hoge" size="lg" variant="outline-primary">
      Post!!<font-awesome-icon icon="coffee" />
    </b-button>
    <div v-if="res != null">
      {{ res }}
    </div>
  </div>
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
        _self.res = response;
        console.log(_self.res);
      });
    },
  },
};
</script>
