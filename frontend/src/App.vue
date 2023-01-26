<template>
    <div class="container" role="navigation" aria-label="main navigation">
      <nav class="navbar is-spaced">
        <div class="navbar-brand">
          <a class="navbar-item" href="/">
            <svg width="112" height="28" viewBox="0 0 60 63" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M2.61331 21.073L30.1883 2L58.0418 21.073M11.2479 26.2386L30.1883 14.0531L48.1538 26.2386M19.7432 31.6692L30.1883 24.1194L40.0763 31.6692" stroke="#EA6526" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M2 22.7761L27.9974 40.5367L55.2818 22.7761" stroke="#B24511" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M2.12871 32.9434L28.1261 50.704L50.0051 36.1609" stroke="#B24511" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M2.64349 43.2394L28.6409 61L43.184 51.0901" stroke="#B24511" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </a>
        </div>
        <div class="navbar-menu">
          <div class="navbar-start">
            <router-link to="/" class="item">Home</router-link>
            <router-link to="/products" class="item">Products</router-link>
            <router-link to="/about" class="item">About</router-link>
          </div>
          <div class="navbar-end">
            <router-link to="/login" class="login">Login</router-link>
            <router-link to="/register" class="login">Register</router-link>
          </div>
        </div>
      </nav>
      <router-view />
    </div>
</template>

<style>
  .item{
    color: #171926;
    padding: 0.5rem 1rem;
    border-radius: 6px;
  }
  .item:hover{
    color: #EA6526;
    border-radius: 6px;
    background-color: #F0F1F2;
  }
  *{
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }
  .container{
    width: 100%;
    height: 100vh;
    background-color: #F0F1F2;
  }
  .login{
    background-color: #EA6526;
    color: #F0F1F2;
    border-radius: 6px;
    padding: 0.5rem 1rem;
  }
</style>
