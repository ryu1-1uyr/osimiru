import axios from 'axios'

class DogApi {
  constructor() {
    this.apiBase = 'http://localhost:8080/';
  }

  breeds() {
    return axios.get(`${this.apiBase}`)
      .then(json =>json.data)
      .catch(e => ({ error: e }));
  }

}

const dogApi = new DogApi();

export default dogApi;
