<script lang="ts">

import Fa from 'svelte-fa'
import { faMagnifyingGlass, faCopy, faCircleExclamation } from '@fortawesome/free-solid-svg-icons'
import { copyText } from 'svelte-copy'

let url: string = ''
let success: boolean = false
let fail: boolean = false
let short: string = ''
let msg: string = ''

const submit = async (e: Event) => {
  e.preventDefault()

  if (url === '') {
    msg = 'URL을 입력해주세요'
    fail = true
    setTimeout(() => fail = false, 3000)
    return
  }

  const res = await fetch(`/api`, {
    method: 'POST',
    body: JSON.stringify({ url })
  })
  const data = await res.json()
  copyText(`http://localhost:5173/${data.key}`)

  msg = '클립보드에 복사했어요!'
  success = true
  setTimeout(() => success = false, 3000)
}
</script>

<main>
  <div class="success0" class:success1={success}>
    <div class="successText">
      <Fa icon={faCopy} />
      {msg}
    </div>
  </div>

  <div class="fail0" class:fail1={fail}>
    <div class="failText">
      <Fa icon={faCircleExclamation} />
      {msg}
    </div>
  </div>

  <div class="container">
    <div class="title">Short</div>

    <form on:submit={submit} class="form">
      <input bind:value={url} placeholder="Type your URL" type="text">
      <button on:click={submit}><Fa icon={faMagnifyingGlass} /></button>
    </form>
  </div>
</main>

<style>
  main {
    width: 100vw;
    height: 100vh;
  }

  .success0,
  .fail0 {
    position: fixed;
    display: flex;
    justify-content: center;
    width: 100%;
    transform: translateY(-50px);
    transition: transform 1s;
    
  }

  .success1,
  .fail1 {
    position: fixed;
    width: 100%;
    display: flex;
    justify-content: center;
    transform: translateY(40px);
    transition: transform 1s;
  }

  .successText {
    background: white;
    backdrop-filter: blur(4px);
    color: rgb(0, 128, 255);
    font-weight: bold;
    font-size: 16px;
    padding: 10px 12px;
    border-radius: 20px;
  }

  .failText {
    background: white;
    backdrop-filter: blur(4px);
    color: rgb(255, 66, 84);
    font-weight: bold;
    font-size: 16px;
    padding: 10px 12px;
    border-radius: 20px;
  }

  .container {
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background: rgb(0,100,255);
    background: linear-gradient(128deg, rgba(0,100,255,1) 0%, rgba(0,182,255,1) 100%);
  }
  
  .title {
    color: white;
    font-size: 35px;
    font-weight: 700;
    margin-top: -48px;
  }

  .form {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: white;
    width: 280px;
    padding: 5px 6px;
    border-radius: 30px;
    box-shadow: 0px 10px 15px -3px rgba(0,0,0,0.1);
  }

  .form > input {
    outline: none;
    border: none;
    font-size: 18px;
    background: transparent;
    color: rgb(0, 128, 255);
    margin-left: 10px;
  }

  .form > button {
    position: relative;
    outline: none;
    border: none;
    color: white;
    background: rgb(0,100,255);
    background: linear-gradient(128deg, rgba(0,100,255,1) 0%, rgba(0,182,255,1) 100%);
    width: 40px;
    height: 40px;
    /* padding: 10px 10px; */
    font-weight: bold;
    font-size: 18px;
    border-radius: 50%;
    cursor: pointer;
  }
</style>