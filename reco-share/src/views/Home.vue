<template>
  <div class="container">
    <b-row>
      <b-col xl="2"></b-col>
      <b-col cols="12" xl="7" class="home">
        <div class="mt-4"></div>
        <b-card tag="article" class="mt-4 mb-3">
          <b-carousel
            id="carousel-fade"
            style="text-shadow: 0px 0px 2px #000"
            fade
            indicators
            img-width="1024"
            img-height="480"
          >
            <b-carousel-slide
              caption="reco×share"
              img-src="https://picsum.photos/1024/480/?image=10"
            ></b-carousel-slide>
            <b-carousel-slide
              caption="reco×share"
              img-src="https://picsum.photos/1024/480/?image=12"
            ></b-carousel-slide>
            <b-carousel-slide
              caption="reco×share"
              img-src="https://picsum.photos/1024/480/?image=22"
            ></b-carousel-slide>
          </b-carousel>
        </b-card>
        <b-card tag="article" class="mb-4 pd-40">
          <h3>
            Let's go on a journey to find the <span class="e5b848">M</span>usic!
          </h3>
          <p class="mb-5">
            好きな音楽・レコードを見つける<br />
            旅に出ましょう<font-awesome-icon icon="plane" class="ml-1" />
          </p>
          <label class="sr-only" for="inline-form-input-name">Name</label>
          <b-row>
            <b-col xl="2"></b-col>
            <b-col xl="8" cols="12">
              <b-input-group prepend="@">
                <b-input
                  id="inline-form-input-name"
                  placeholder="ユーザ名、またはメールアドレス"
                  aria-describedby="password-help-block"
                  v-model="users.user"
                ></b-input>
              </b-input-group>

              <b-input-group class="mt-2">
                <b-input
                  id="text-password"
                  type="password"
                  placeholder="パスワード"
                  v-model="users.pass"
                ></b-input>
              </b-input-group>
            </b-col>
            <b-col xl="2"></b-col>
          </b-row>
          <br />

          <b-button class="mt-3" href="#" :disabled="isInputForm"
            >Let's go!!<font-awesome-icon icon="ship" class="ml-1"
          /></b-button>

          <hr />
          <p>アカウントをお持ちでないですか？</p>
          <b-button
            >おためしログイン<font-awesome-icon icon="sign-in-alt"
          /></b-button>
          <br />
          <b-button class="mt-1 mb-5"
            >新規登録<font-awesome-icon icon="user-plus" class="ml-1"
          /></b-button>
        </b-card>
      </b-col>
      <b-col xl="3"></b-col>
    </b-row>
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
      users: {
        user: null,
        pass: null,
      },
    };
  },
  created() {},
  computed: {
    isInputForm() {
      return (
        this.users.user == null ||
        this.users.pass == null ||
        this.users.user == "" ||
        this.users.pass == ""
      );
    },
  },

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

<style scoped>
.btn {
  color: black;
  background-color: white;
  box-shadow: 2px 2px 4px #cccccc;
  width: 170px;
}
.btn:hover {
  background-color: #e5b848;
  color: black;
}
.card {
  box-shadow: 2px 2px 4px #cccccc;
}
.pd-40 {
  padding: 40px;
}
.e5b848 {
  color: #e5b848;
}
.gothic {
  font-family: "游ゴシック Medium", "Yu Gothic Medium", "游ゴシック体",
    "YuGothic", "ヒラギノ角ゴ ProN W3", "Hiragino Kaku Gothic ProN", "メイリオ",
    "Meiryo", "verdana", sans-serif;
}
.home {
  background-color: #f4f5f7;
  box-shadow: 1px 1px 1px 1px #cccccc;
}
.title {
  font-weight: 800;
}
</style>
