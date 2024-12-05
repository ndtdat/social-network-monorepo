import * as grpcWeb from 'grpc-web';

import * as user_rpc_register_pb from '../user/rpc/register_pb';


export class UserClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  register(
    request: user_rpc_register_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_rpc_register_pb.RegisterReply) => void
  ): grpcWeb.ClientReadableStream<user_rpc_register_pb.RegisterReply>;

}

export class UserPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  register(
    request: user_rpc_register_pb.RegisterRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<user_rpc_register_pb.RegisterReply>;

}

