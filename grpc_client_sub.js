const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = __dirname + '/subtraction/subtraction.proto';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
});
const packageDescriptor = grpc.loadPackageDefinition(packageDefinition);
const subtraction = packageDescriptor.subtraction;

const stub = new subtraction.Subtraction(
  'localhost:10000',
  grpc.credentials.createInsecure()
);

const minusRequest = {
  number: process.argv[2] || 0,
  anotherNumber: process.argv[3] || 0
};

stub.Minus(minusRequest, function(err, response) {
  if (err) {
    console.error(err);
  } else {
    console.log(
      `The subtract of ${minusRequest.number} and ${minusRequest.anotherNumber} is ${
        response.subtract
      }`
    );
  }
});

