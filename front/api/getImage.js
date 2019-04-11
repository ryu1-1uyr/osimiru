import axios from 'axios'

class ImageAPI {
  constructor() {
    this.apiBase = 'http://localhost:8080/roa';
  }

  initialGet() {
    return axios.get(this.apiBase)
      .then(json =>json.data)
      .catch(e => ({ error: e }));
  }

  osiget(url){
    return axios.get(url)
      .then(json =>json.data)
      .catch(e => ({ error: e }));
  }


}

const imageAPI = new ImageAPI();
export default imageAPI;
