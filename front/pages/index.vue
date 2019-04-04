<template>
  <section class="container">
    <div class="columns is-multiline">

      <!-- v-for で breed_list からループ出力　mapstateで名前をさぶすてーとと揃えた時は直接呼べるっぽい -->
      <div v-for="(i) in breed_list.URL" v-bind:key='i'>
        <img :src="i" >
      </div>

      <!--<p>{{breed_list.URL}}</p><img :src="breed_list.image" >-->
    </div>
  </section>
</template>


<script>
  import dogApi from '@/api/getImage'
  import {mapState} from 'vuex'

  export default {
    async fetch({store}) {
      let json = await dogApi.breeds();
      store.commit('breed_list_update', json)//axiosで取ってきたjsonをデータストアにのっける
    },
    computed: mapState(['breed_list']),
    // 複雑なロジックには算出プロパティを使うべきらしいのでmethodじゃなくてcomputedつかう(https://jp.vuejs.org/v2/guide/computed.html)
    // mapState ヘルパー vuexで作ったデータストアなのでvuexのめそっどから呼んでるよ
    //ステートサブツリーの名前と同じ場合は、文字列配列を mapState に渡すこともできますらしいです
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

