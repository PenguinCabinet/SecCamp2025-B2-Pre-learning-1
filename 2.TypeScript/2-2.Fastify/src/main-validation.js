const fastify = require('fastify')({
    logger: true
})

/**
 * @type {import('fastify').RouteShorthandOptions}
 * @const
 */
const opts = {
    schema: {
        body: {
            type: 'object',
            required: ['someKey', 'someOtherKey'],
            properties: {
                someKey: { type: 'string' },
                someOtherKey: { type: 'number' }
            }
        }
    }
}

fastify.post('/', opts, async (request, reply) => {
    reply.send(request.body)
})

// Run the server!
fastify.listen({ port: 3000 }, function (err, address) {
    if (err) {
        fastify.log.error(err)
        process.exit(1)
    }
    // Server is now listening on ${address}
})