<template>
  <div id="app">
    <button @click="sendRequest">リクエスト送信</button>
    <h1>取得結果</h1>
    <p>GET(パラメータ無):<br/><strong>{{title}}</strong></p>
    <p>GET(パラメータ有):<br/><strong>{{name1}}</strong></p>
    <p>POST(form-urlencoded):<br/><strong>{{name2}}</strong></p>
    <p>POST(JSONデータ):<br/><strong>{{company}}</strong></p>
  </div>
</template>

<script>
import axios from "axios"
//
export default {
  name: 'App',
  data:()=>{
    return{
      title:"",
      name1:"",
      name2:"",
      company:""
    }
  },
  methods: {
    sendRequest: async function(){
      //パラメータ無しでGETリクエスト
      const getRequestNoParam = await axios.get("http://localhost:8000/getTitle")

      //パラメータ付きでGETリクエスト
      const getRequest = await axios.get("http://localhost:8000/getName/ひふみん")

      //application/x-www-form-urlencodedでデータを送信
      const params = new URLSearchParams();
      params.append("name","青葉");
      const postRequest = await axios.post("http://localhost:8000/postName",params)

      //JSONデータを送信(axiosはデフォルトでJSONを送信)
      const jsonPostRequest = await axios.post("http://localhost:8000/postCompany", {
        company: "Eagle Jump",
        works: "PECO"
      });

      //取得結果をviewに反映
      this.title = getRequestNoParam.data
      this.name1 = getRequest.data
      this.name2 = postRequest.data
      this.company = jsonPostRequest.data
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
