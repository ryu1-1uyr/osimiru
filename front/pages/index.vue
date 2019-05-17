<template>

<div>

  <dvi style="padding-top: 20px;margin: auto">
    <a @click="setOsi('roa')">ろあ</a> <a @click="setOsi('toko')">いぬい</a> <a @click="setOsi('gibara')">えら</a>
  </dvi>

  <section class="container">

    <div class="columns is-multiline" style="margin-top: 40px;">

      <div v-for="(i) in url_list.URL" v-bind:key='i'>
          <img :src="shaping(i)" >
      </div>

    </div>
  </section>
</div>
</template>


<script>
  import getImageAPI from '@/api/getImage'
  import {mapState} from 'vuex'

  const elements = "https://pbs.twimg.com/media"

  export default {
    data () {
      return {
        url : "http://localhost:8080/",
        osiName: ""
      }
    },
    async fetch({store}) {
      let json = await getImageAPI.initialGet();
      store.commit('url_list_update', json)
    },
    computed: mapState(['url_list']),

    methods: {
      shaping (url) {
        if(url.indexOf(elements) === 0){
          return url
        }
      },

      async switchImages () {

        let json = await getImageAPI.osiget(this.url + this.osiName);
        this.$store.commit('url_list_update', json)

      },

      setOsi (name) {
        this.osiName = name;
        console.log(name)

        this.switchImages()
      }

    }

  }
</script>

<style>
  .container {
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
  }

  .title {
    font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
    display: block;
    font-weight: 300;
    font-size: 100px;
    color: #35495e;
    letter-spacing: 1px;
  }

  .subtitle {
    font-weight: 300;
    font-size: 42px;
    color: #526488;
    word-spacing: 5px;
    padding-bottom: 15px;
  }

  .links {
    padding-top: 15px;
  }
</style>

