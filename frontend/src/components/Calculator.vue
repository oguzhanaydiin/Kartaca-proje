<template>

<div class="Calculator">
    
    
<div data-spy="scroll" data-target=".site-navbar-target" data-offset="300">

  
  
  <div class="site-wrap"  id="home-section">
      
    <header class="site-navbar js-sticky-header site-navbar-target" role="banner">

      <div class="container">
        <div class="row align-items-center position-relative">
          
            
            <div class="site-logo">
              <a href="index.html" class="text-black"><span class="text-primary"/>Oguzhanaydin</a>
            </div>
            
            <div class="col-4">
              <nav class="site-navigation text-right ml-auto " role="navigation">

                <ul class="site-menu main-menu js-clone-nav ml-auto d-none d-lg-block">
                <li><a href="#home-section" class="nav-link"><b>Upwards</b></a></li>

                </ul>
              </nav>
          
            </div>

            <div class="col-8">
              <nav class="site-navigation text-right ml-auto " role="navigation">

                <ul class="site-menu main-menu js-clone-nav ml-auto d-none d-lg-block">
                
                <li><input type="text" class="btn btn-block border-gray-4 py-1 px-2" style="border-color: coral" placeholder="E-mail"></li>
                <li><input type="text" class="btn btn-block border-gray-4 py-1 px-2" style="border-color: coral" placeholder="Password"></li>
                <li><input type="submit" class="btn btn-block btn-primary text-white py-1 px-5" value="Login"></li>

                </ul>
              </nav>
          
            </div>

          <div class="toggle-button d-inline-block d-lg-none"><a href="#" class="site-menu-toggle py-5 js-menu-toggle text-black"><span class="icon-menu h3"></span></a></div>

        </div>
      </div>
      
    </header>

    <div class="site-section" id="team-section">
      <div class="container">
        <div class="row mb-5 justify-content-center">
          <div class="col-md-7 text-center">
            <div class="block-heading-1" data-aos="fade-up" data-aos-delay="">
              <h2>Latest Posts</h2>
             </div>
          </div>
        </div>
        <div class="row">

          <div class="col-lg-12 col-md-6 mb-4" data-aos="fade-up">

            <div class="block-team-member-1 text-center rounded">
              <h2 class="font-size-20">Jean Smith</h2>
              <p class="d-block font-black-5 letter-spacing-1  font-size-20 mb-3">Mins asdklasjkldna askldj lkasjd lkas dlkasdj lkadfasdf asdfşlmasdlf sdlkfmnaklsdn asdfnasdkf nakdsf kasdfnkasndfking </p>
              
            </div>

          </div>

        </div>
      </div>
    </div>


    <div v-if="render!='dont'" class="site-section bg-light" id="register-section">
      <div class="container">
        <div class="row">
          <div class="col-12 text-center mb-5" data-aos="fade-up" data-aos-delay="">
            <div class="block-heading-1">
              <span>You're not logged in!</span>
              <h2>Register</h2>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="col-lg-12 mb-5" data-aos="fade-up" data-aos-delay="100">

            <form @submit.prevent>
              <div class="form-group row">
                <div class="col-md-6 mb-4 mb-lg-0">
                  <input name="firstname" v-model="firstname" type="text" class="form-control" placeholder="First name" required>
                </div>
                <div class="col-md-6">
                  <input name="lastname" v-model="lastname" type="text" class="form-control" placeholder="Last name" required>
                </div>
              </div>

              <div class="form-group row">
                <div class="col-md-12">
                  <input name="mail" v-model="mail" type="text" class="form-control" placeholder="Email address" required>
                </div>
              </div>

              <div class="form-group row">
                <div class="col-md-12">
                  <input name="password" v-model="password" type="password" class="form-control" placeholder="Password" required>
                </div>
              </div>

              <div class="form-group row">
                <div class="col-md-12">
                  <input name="repassword" v-model="repassword" type="password" class="form-control" placeholder="Re-Password" required>
                </div>
              </div>
              
              <div class="form-group row">
                <div class="col-md-12 ml-auto">
                  <button variant="primary" v-on:click="postreq()" class="btn btn-block btn-primary text-white py-3 px-5" value="Submit">Register</button>
                </div>
              </div>
            </form>

            

          </div>
        </div>
      </div>
    </div>

    <div v-if="render=='do' || render =='dont'" class="col-md-12">
              <div style="background-color:orange"><label class="label">{{ errormessage }}</label></div>
    </div>

    <footer class="site-footer">
      <div class="container">
        <div class="row">
          <div class="col-md-6">
            <div class="row">
              <div class="col-md-10">
                <h2 class="footer-heading mb-4">About Us</h2>
                <p>This website is for challenge. You can register, login and post things. Also you can click the posts for showing much info. Have fun!</p>
              </div>
              <div class="col-md-2 ml-auto ">
                <h2 class="footer-heading mb-4">Features</h2>
                <ul class="main-menu">
                  <li><a href="#register-section" >Register</a></li>
                  <li><a href="#home-section">Home</a></li>
                </ul>
              </div>
              
            </div>
          </div>
        </div>
      </div>
    </footer>

  </div>
</div>
</div>

</template>


<style>
  @import 'css/style.css';
  @import 'css/bootstrap.min.css';
  @import 'https://fonts.googleapis.com/css?family=Rubik:300,400,700|Oswald:400,700';
</style>

<script>

import axios from 'axios';

export default {
  name: 'Calculator',
  
  data: function() {
    return {
      firstname: "", lastname: "", password: "", repassword: "", 
      mail: "", errorcode: "", errormessage: "", render:"", postData : [], postItems: { firstname:null, lastname:null, date:null  }
    }
  },

  methods: {
    postreq: function() {
      var data = {"firstname": this.firstname, "lastname": this.lastname, "mail": this.mail, "password": this.password, "repassword":this.repassword}

      /*eslint-disable*/
      console.log(data) 
      /*eslint-enable*/

      axios({ method: "POST", url: "http://127.0.0.1:8090/calc", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
          // this.response = result.data;
          this.errorcode = result.data['errorcode']
          this.errormessage = result.data['errormessage']
          if(this.errorcode == 0){
             this.render = "dont"
          }else{
            this.render = "do"
          }
          /*eslint-disable*/
          console.log(result.data) 
          /*eslint-enable*/

        }).catch( error => {
            /*eslint-disable*/
            console.error(error);
            /*eslint-enable*/
      });
    },

    //getposts: function() {
    //  var data = {postData: [], postItems: null}

      /*eslint-disable*/
    //  console.log(data) 
      /*eslint-enable*/

    //  axios({ method: "POST", url: "http://127.0.0.1:8090/forposts", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
            /*eslint-disable*/
    //        console.log(result.data) 
            /*eslint-enable*/
    //      this.postData = result

    //    }).catch( error => {
            /*eslint-disable*/
    //        console.error(error);
            /*eslint-enable*/
    //  });
    //}

  },
  //beforeMount(){
  //  this.getposts()
    
  //}
}
</script>

