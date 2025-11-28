import { Page } from "@/model/base";
import { QueryRequest, RecordItem } from "@/model/record";
import { ApiResponse, post } from "@/utils/request";

export async function listRecords(query: QueryRequest): Promise<ApiResponse<Page<RecordItem>>> {
  return await post<Page<RecordItem>>('/record', query)
}
