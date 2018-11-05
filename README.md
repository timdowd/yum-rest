# YumREST
Examples of Yum RESTful interactions


### Topics:

- ID, Id, or _id
  - use {thing}Id in protobuf vs just id 
  - thingId

- Key for update
  - in header or body
  - can it be changed (if not uuid)
  - RFC 72314.3.4 PUT:
    - PUT should have same URI as GET
  
- Quiet failure for delete?
  - 200 with statusCode and status
  - 204 with no body
  - 404 for failure

- Scrub _id from PUT swagger entry?

- Always use accessors

- Location header from POST
  - returns _id
  - do we need the header itself?
  - Location vs Grpc-Metadata-Location

- PUT as upsert?
  - 200 success
  - 201 upsert
  - 204 no content

- PATCH: replaces part of an entity
  - PUT should be replacing whole thing
  - we can probably just stick to PUT

- CORS
  - middleware in phdmw
  - whitelist ui urls
 
- Logging
  - log beginning and end of handler
  - phdmw automatically logs end of handler
  - 

- Tracing
  - Trace any call to an external service
    - datastore
    - other rpc
    - external api
  - use StartSpan, not NewSpan ie. GetThing

- ConvIDs