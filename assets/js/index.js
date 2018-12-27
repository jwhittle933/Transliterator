
const transliterator = new Vue({
    el: "#app",
    data() {
      return {
        testText: "This is a test",
        greek: false,
        langOptions: ["Greek", "Hebrew", "Syriac", "Aramaic"],
        currentLang: ""
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
      },
      langDown(){
        var currentIndex = this.langOptions.findIndex(x => x === this.currentLang)
        if (currentIndex === 0){
          this.currentLang = this.langOptions[this.langOptions.length - 1]
        } else {
          this.currentLang = this.langOptions[currentIndex - 1]
        }
      }
    }
  })