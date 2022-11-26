# Proxy

> Proxy is a structural design pattern that lets you provide a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.

## assumption

imagine you are building a vpn software to help user accessing some address that they can not access directly.
you need to use a proxy which can visit those address, when user requesting those address, proxy will access those address first, then cache the result, and send back to user
if request address has already cached, then you can reply to user directly