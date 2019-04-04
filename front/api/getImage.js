// import axios from 'axios'
//
// class GetImageApi {
//   constructor() {
//     this.apiBase = 'https://pokeapi.co/api/v2/pokemon/1';
//   }
//
//   osiNoGazou() {
//     return axios.get(`${this.apiBase}`)
//       .then(json => {
//         return json;
//       })
//       .catch(e => ({ error: e }));
//   }
//   //今までみたいに1つのおっきいデータを叩くんじゃなくて、必要な箇所箇所をaxios.getで取ってくる書き方ができる
//   //["hoge"]["huga"][0]みたいな感じで深くまで自前でアクセスしないでもいいんやで…
// }
//
// const getimageAPI = new GetImageApi();
//
// export default getimageAPI;

import axios from 'axios'
const elements = "https://pbs.twimg.com/media";

class DogApi {
  constructor() {
    this.apiBase = 'http://localhost:8080/';
  }

  breeds() {
    return axios.get(`${this.apiBase}`)
      .then(json =>json.data)
      // .then(url => url.indexOf(elements) === 0 ? i : ''  )
      .catch(e => ({ error: e }));
  }

  //今までみたいに1つのおっきいデータを叩くんじゃなくて、必要な箇所箇所をaxios.getで取ってくる書き方ができる
  //["hoge"]["huga"][0]みたいな感じで深くまで自前でアクセスしないでもいいんやで…
}

const dogApi = new DogApi();

export default dogApi;
