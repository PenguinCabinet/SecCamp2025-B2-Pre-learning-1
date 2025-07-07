const fastify = require('fastify')({
    logger: true
})

/**
 * @type {import('fastify').RouteShorthandOptions}
 * @const
 */
const opts = {
    schema: {
        response: {
            200: {
                type: 'object',
                properties: {
                    hello: { type: 'string' }
                }
            }
        }
    }
}

fastify.get('/', opts, async (request, reply) => {
    return { hello: 'world' }
})

// Run the server!
fastify.listen({ port: 3000 }, function (err, address) {
    if (err) {
        fastify.log.error(err)
        process.exit(1)
    }
    // Server is now listening on ${address}
})