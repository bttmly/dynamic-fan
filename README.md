# turbofan

A little example of using `reflect` to receive on an arbitrary-sized slice of channels. Whenever a value is received on any channel, it is sent to every other channel.

### `New(chans ...chan bool)`
Returns a new `Turbofan` listening to all of `chans`

### `t.Broadcast(b bool)`
Broadcasts a value to all channels.

### `t.Close()`
Closes all channels.

To do... special handling on on `Broadcast()` after `Close()`?

