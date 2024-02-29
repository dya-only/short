<script lang="ts">
import { Router, Route } from 'svelte-routing'
import Short from './pages/Short.svelte'
import Redirect from './lib/Redirect.svelte'
</script>

<Router>
  <Route path="/" component={Short} />
  <Route path="/:key" component={Redirect} />
</Router>
