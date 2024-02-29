<script lang="ts">
import Fa from 'svelte-fa'
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons'

let url: string = ''

const submit = async (e: Event) => {
  e.preventDefault()

  const res = await fetch(`/api`, {
    method: 'POST',
    body: JSON.stringify({ url })
  })
  const data = await res.json()

  console.log(data.key)
}
</script>

<main>


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
    display: flex;
    flex-direction: column;
    align-items: center;
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