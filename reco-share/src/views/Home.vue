<template>
  <div class="container">
    <b-row>
      <b-col xl="2"></b-col>
      <b-col cols="12" xl="7" class="home">
        <div class="mt-4"></div>
        <b-card tag="article" class="mt-4 mb-4">
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
        <b-card tag="article" class="mt-4 mb-4 pd-40">
          <h3>
            Let's go on a journey to find the <span class="e5b848">M</span>usic!
          </h3>
          <p class="mb-5">
            好きな音楽・レコードを見つける<br />
            旅に出ましょう<font-awesome-icon icon="plane" class="ml-1" />
          </p>
          <p v-if="loginFailure" style="color:red">
            ユーザ名、またはパスワードが間違っています。
          </p>
          <label class="sr-only" for="inline-form-input-name">Name</label>
          <b-row>
            <b-col xl="2"></b-col>
            <b-col xl="8" cols="12">
              <b-input-group prepend="@">
                <b-input
                  id="inline-form-input-name"
                  placeholder="ユーザ名、またはメールアドレス"
                  v-model="users.userNm"
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

          <b-button
            class="mt-3 mb-3"
            href="#"
            :disabled="isInputForm"
            @click="login(users)"
            >Let's go!!<font-awesome-icon icon="ship" class="ml-1"
          /></b-button>

          <hr />
          <p class="mt-5">アカウントをお持ちでないですか？</p>
          <b-button @click="guestLogin()"
            >おためしログイン<font-awesome-icon icon="sign-in-alt"
          /></b-button>
          <br />
          <b-button
            v-b-modal.modal-1
            @click="showRegistModal()"
            class="mt-1 mb-5"
            ref="btnshow"
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
import Eventbus from "../eventbus";
export default {
  name: "Home",
  components: {},
  data() {
    return {
      res: null,
      users: {
        userNm: null,
        pass: null,
      },
      loginFailure: false,
    };
  },
  created() {},
  computed: {
    isInputForm() {
      return (
        this.users.userNm == null ||
        this.users.pass == null ||
        this.users.userNm == "" ||
        this.users.pass == ""
      );
    },
  },

  methods: {
    guestLogin() {
      let _self = this;
      _self.$axios.post("http://localhost:8082/guestlogin").then((r) => {
        if (r) {
          let now = new Date();
          now.setMinutes(now.getMinutes() + 10); // 10分後に設定。
          document.cookie =
            "id=" +
            encodeURIComponent(r.data.UserID) +
            ";expires=" +
            now.toUTCString();
          location.href = "http://localhost:8080/timeline";
        }
      });
    },
    login(users) {
      let _self = this;
      _self.loginFailure = false;
      _self.$axios.post("http://localhost:8082/login", users).then((r) => {
        // ログイン成功時は、タイムライン画面に遷移する。
        if (r.data.resultCode == "00") {
          let now = new Date();
          now.setMinutes(now.getMinutes() + 10); // 10分後に設定。
          document.cookie =
            "id=" +
            encodeURIComponent(r.data.UserID) +
            ";expires=" +
            now.toUTCString();
          location.href = "http://localhost:8080/timeline";
        } else {
          // ログイン失敗時は、ログイン失敗の文言を画面に表示する。
          _self.loginFailure = true;
        }
      });
    },
    showRegistModal() {
      Eventbus.$emit("open-registration-modal");
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
.pd-40 {
  padding: 40px;
}
.title {
  font-weight: 800;
}
</style>
