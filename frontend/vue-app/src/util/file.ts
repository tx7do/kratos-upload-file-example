import axios, { AxiosProgressEvent } from 'axios';
import { message } from 'ant-design-vue';
import { UploadProgressEvent, UploadRequestError } from 'ant-design-vue/es/vc-upload/interface';
import { API_URL } from '@/util/const';
import { UploadMethod } from '@/rpc/api/file/service/v1/file.pb';

export function makeFormData(file: File, data: object) {
  const formData = new FormData();

  if (data !== null && data !== undefined) {
    Object.entries(data)
      .forEach(([k, v]) => {
        formData.append(k, v);
      });
  }

  formData.append('file', file);

  return formData;
}

export function putFile(
  url: string,
  file: File,
  data: object,
  onProgress?: (event: UploadProgressEvent) => void,
  onSuccess?: (body: any, xhr?: XMLHttpRequest) => void,
  onError?: (event: UploadRequestError | ProgressEvent, body?: never) => void,
) {
  const formData = makeFormData(file, data);

  axios
    .put(url, formData, {
      headers: {
        'Content-Type': file.type,
        // 'Access-Control-Allow-Origin': '*',
      },
      onUploadProgress: (progressEvent: AxiosProgressEvent) => {
        const percentCompleted = Math.round(
          (progressEvent.loaded * 100) / (progressEvent.total ?? 0),
        );
        onProgress?.({ percent: percentCompleted });
      },
    })
    .then(async function (response) {
      if (response.status === 200 || response.status === 204) {
        onSuccess?.(response.data);
        await message.info(`${file.name} upload success!`);
        console.info(`${file.name} upload success!`);
      } else {
        onError?.(response?.data?.message ?? response.statusText);
        await message.error(`${file.name} upload failed!`);
        console.error(`${file.name} upload failed!`);
      }
    })
    .catch(async function (error) {
      onError?.(error?.response?.data?.message ?? error.message);
      await message.error(`${file.name} upload exception!`, error);
      console.error(`${file.name} upload exception!`, error);
    });
}

export function postFile(
  url: string,
  file: File,
  data: object,
  needMakeFormData: boolean,
  onProgress?: (event: UploadProgressEvent) => void,
  onSuccess?: (body: any, xhr?: XMLHttpRequest) => void,
  onError?: (event: UploadRequestError | ProgressEvent, body?: never) => void,
) {
  const formData = makeFormData(file, data);

  axios
    .post(url, formData, {
      onUploadProgress: (progressEvent: AxiosProgressEvent) => {
        const percentCompleted = Math.round(
          (progressEvent.loaded * 100) / (progressEvent.total ?? 0),
        );
        onProgress?.({ percent: percentCompleted });
      },
    })
    .then(async function (response) {
      if (response.status === 200 || response.status === 204) {
        onSuccess?.(response.data);
        await message.info(`${file.name} upload success!`);
        console.info(`${file.name} upload success!`);
      } else {
        onError?.(response?.data?.message ?? response.statusText);
        await message.error(`${file.name} upload failed!`);
        console.error(`${file.name} upload failed!`);
      }
    })
    .catch(async function (error) {
      await message.error(`${file.name} upload exception!`, error);
      console.error(`${file.name} upload exception!`, error);
      onError?.(error?.response?.data?.message ?? error.message);
    });
}

export function retrievePutUrl(file: File, cb: (file: File, url: string) => void) {
  const url = `${API_URL}/file:upload-url`;

  const data = {
    method: UploadMethod.Put,
    contentType: file.type,
    fileName: file.name,
    bucketName: 'files'
  };

  axios.post(url, data)
    .then(function (response) {
      cb(file, response.data.uploadUrl);
    })
    .catch(function (error) {
      console.error(error);
    });
}

export function retrievePostUrl(file: File, cb: (file: File, url: string, data: object) => void) {
  const url = `${API_URL}/file:upload-url`;

  const data = {
    method: UploadMethod.Post,
    contentType: file.type,
    fileName: file.name,
    bucketName: 'files'
  };

  axios.post(url, data)
    .then(function (response) {
      cb(file, response.data.uploadUrl, response.data.formData);
    })
    .catch(function (error) {
      console.error(error);
    });
}
