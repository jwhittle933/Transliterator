<template>
    <div class="jumbotron container mt-5 mx-auto w-75">
      <h1 class="text-center mb-5 font-main font-xl">Transliterator</h1>
      <div class="d-flex justify-content-between p-3">
        <div class="form-group mb-2 w-50">
          <span class="ml-1 font-lg pointer">
            <i class="fas fa-angle-up" @click="langUp"></i>
          </span>
          <div class="bg-dark mt-3 rounded width-md">
            <p class="text-light ml-1 font-lg font-main">{{ currentLang }}</p>
          </div>
          <span class="ml-1 font-lg pointer">
            <i class="fas fa-angle-down" @click="langDown"></i>
          </span>
        </div>
        <div class ="form-group mx-sm-3 mb-2">
          <textarea name="text" type="text" cols="50" rows="10" :dir="direction" v-model="text" class="form-control bg-dark text-white" placeholder="..."></textarea>
        </div>
      </div>
      <input type="hidden" name="lang" :value="currentLang"/>
      <button type="submit" class="btn btn-block btn-outline-primary mt-5" @click.prevent="submit">Transliterate!</button>
      <div>
        <response-text
          :response="responseText"
          v-if="response"
        ></response-text>
      </div>
    </div>
</template>

<script>
import axios from 'axios'
import ResponseText from './ResponseText.vue'

export default {
    name: 'app',
    components: {
      ResponseText
    },
    data() {
      return {
        langOptions: ["Greek", "Hebrew", "Syriac", "Aramaic"],
        currentLang: "",
        direction: "ltr",
        response: true,
        responseText: "",
        text: "",
        transText: ""
      }
    },
    created() {
        this.currentLang = this.langOptions[0]
    },
    methods: {
      langUp() {
        var currentIndex = this.langOptions.findIndex(x => x === this.currentLang)
        if (currentIndex === this.langOptions.length - 1){
          this.currentLang = this.langOptions[0]
        } else {
          this.currentLang = this.langOptions[currentIndex + 1]
        }
        this.currentLang !== "Greek" ? this.direction = "rtl" : this.direction = "ltr"
      },
      langDown(){
        var currentIndex = this.langOptions.findIndex(x => x === this.currentLang)
        if (currentIndex === 0){
          this.currentLang = this.langOptions[this.langOptions.length - 1]
        } else {
          this.currentLang = this.langOptions[currentIndex - 1]
        }
        this.currentLang !== "Greek" ? this.direction = "rtl" : this.direction = "ltr"
      },
      submit(){
            axios.post("/", { params: {
                    language: this.currentLang,
                    text: this.text
                }
            }).then( res => {
                this.transText = res.data
            }).catch( e => { this.transText = "Sorry, there was an error on our side: " + e})
      }
    }
}
</script>

<style scoped>
.font-main{
    font-family: Nunito;
}
.font-lg {
    font-size: 3em;
}
.font-xl {
    font-size: 3.5em;
}
.width-md {
    width: 15vw;
}
.pointer {
    cursor: pointer;
}
</style>
