<template>
  <div class="home">
    <button @click="hoge">hogehoge</button>
    <div v-if="res != null">
      {{ res }}
    </div>
    <b-button size="lg" variant="outline-primary">
      Bootstrap Button
    </b-button>
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
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
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
