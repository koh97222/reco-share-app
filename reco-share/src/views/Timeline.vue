<template>
  <div class="timeline">
    <div class="container">
      <div class="title">
        <h2>FEATURED</h2>
        <hr />
      </div>
      <div class="latest-post title">
        <h2>LATEST POST</h2>
        <hr />
        <template>
          <div class="container">
            <b-row>
              <b-col xl="2"></b-col>
              <b-col xl="8" cols="12">
                <b-card>
                  <b-row>
                    <b-col xl="2" cols="3"><b-badge>title</b-badge></b-col>
                    <b-col xl="10" cols="9"
                      ><h6 class="text-left">
                        <b-link class="text-left"
                          >最近買ったターンテーブル</b-link
                        >
                      </h6></b-col
                    >
                  </b-row>
                  <b-row>
                    <b-col xl="2" cols="3"><b-badge>tag</b-badge></b-col>
                    <b-col xl="10" cols="9"
                      ><h6 class="text-left">
                        <b-link class="text-left">#record </b-link>
                        <b-link class="text-left">#vinyl </b-link>
                        <b-link class="text-left">#tokyo </b-link>
                      </h6></b-col
                    >
                  </b-row>
                  <b-row>
                    <b-col xl="2" cols="3"><b-badge>user</b-badge></b-col>
                    <b-col xl="10" cols="9"
                      ><h6 class="text-left">
                        <b-link class="text-left">@user</b-link>
                      </h6></b-col
                    >
                  </b-row>
                  <hr />
                  <div class="img">
                    image
                  </div>
                  <hr />
                </b-card>
              </b-col>
              <b-col xl="2"></b-col>
            </b-row>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style>
@media only screen and (max-width: 750px) {
  img {
    max-width: 100%;
  }
}
.badge-secondary {
  width: 60px;
}
.timeline {
  background-color: #f2f2f2;
}
.title {
  padding: 15px;
}
</style>
