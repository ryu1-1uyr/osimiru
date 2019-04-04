import axios from 'axios'

class GetImageApi {
  constructor() {
    this.apiBase = 'http://localhost:8080/';
  }

  osiNoGazou() {
    return axios.get(`${this.apiBase}`)
      .then(json => {
        return json.data.message;
      })
      .catch(e => ({ error: e }));
  }
  //今までみたいに1つのおっきいデータを叩くんじゃなくて、必要な箇所箇所をaxios.getで取ってくる書き方ができる
  //["hoge"]["huga"][0]みたいな感じで深くまで自前でアクセスしないでもいいんやで…
}

const getimageAPI = new GetImageApi();

export default getimageAPI;
