package com.atlas.dis;

import java.net.URI;

import com.atlas.dis.constant.RestConstants;
import com.atlas.dis.processor.MonsterDropProcessor;
import com.atlas.shared.rest.RestServerFactory;
import com.atlas.shared.rest.UriBuilder;

import database.PersistenceManager;

public class Server {
   public static void main(String[] args) {
      PersistenceManager.construct("atlas-dis");

      URI uri = UriBuilder.host(RestConstants.SERVICE).uri();
      RestServerFactory.create(uri, "com.atlas.dis.rest");

      MonsterDropProcessor.init();
   }
}
