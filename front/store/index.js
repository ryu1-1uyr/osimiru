import Vuex from 'vuex'

//vuexつかってストアを定義。くろーじゃーっぽく書いてるのは多分外部要因でデータを書き換えさせないため？同一データを参照するためかもしれん(呼ばれるたびに新しくnewサレルノ防いでる？)
const appStore = () => {
  return new Vuex.Store({
    state: {
      images_list: {},
    },
    mutations: {
      //犬種の一覧を取得
      images_list_update(state, payload) {
        state.images_list = {...payload}
      },

    }
  })
};

export default appStore;
