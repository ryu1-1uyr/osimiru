import Vuex from 'vuex'

const appStore = () => {
  return new Vuex.Store({
    state: {
      url_list: {},
    },
    mutations: {
      url_list_update(state, payload) {
        state.url_list = {...payload}
      },

    }
  })
};

export default appStore;
