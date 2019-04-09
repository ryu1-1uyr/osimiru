import axios from 'axios'

class ImageAPI {
  constructor() {
    this.apiBase = 'http://localhost:8080/';
  }

  osiGet() {
    return axios.get(`${this.apiBase}`)
      .then(json =>json.data)
      .catch(e => ({ error: e }));
  }

}

const imageAPI = new ImageAPI();
export default imageAPI;
